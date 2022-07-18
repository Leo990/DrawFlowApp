package services

import (
	"DrawFlowApp/dao"
	"DrawFlowApp/models"

	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os/exec"

	"github.com/dgraph-io/dgo/v2"
	"github.com/dgraph-io/dgo/v2/protos/api"
	"google.golang.org/grpc"
)

type ProgramService struct {
}

func ConnectToDatabase() (*dgo.Txn, *grpc.ClientConn, context.Context) {
	ctx := context.TODO()

	/* Conexion a la base de datos "https://blue-surf-590978.us-east-1.aws.cloud.dgraph.io/graphql" */
	conn, err := grpc.Dial("localhost:9080", grpc.WithInsecure())
	if err != nil {
		log.Fatal("failed to dial ", err)
	}
	//defer conn.Close()
	dgraphClient := dgo.NewDgraphClient(api.NewDgraphClient(conn))

	if err != nil {
		log.Fatal("failed to dial ", err)
	}

	txn := dgraphClient.NewTxn()

	return txn, conn, ctx
}

func (ps ProgramService) Index() (map[string]interface{}, error) {
	txn, _, ctx := ConnectToDatabase()

	q := `{
		programs(func: has(nodes)) {
		uid
		program_name
		}
		}`

	data := map[string]interface{}{}

	res, err := txn.Query(ctx, q)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		json.Unmarshal(res.GetJson(), &data)
	}

	return data, err
}

func (ps ProgramService) Store(new_program []byte) (*api.Response, error) {
	txn, _, ctx := ConnectToDatabase()

	mu := &api.Mutation{
		SetJson: new_program,
	}

	req := &api.Request{CommitNow: true, Mutations: []*api.Mutation{mu}}

	res, err := txn.Do(ctx, req)
	return res, err
}

func (ps ProgramService) Execute(execute_program []byte) (map[string]string, error) {
	var program models.Program
	path := "./files/program.py"

	response := map[string]string{}

	json.Unmarshal(execute_program, &program)

	strProgram := program.WriteProgram()

	dao.WriteProgram(strProgram, path)

	cmd := exec.Command("python", "D:/leona/Documents/Projects/Go/src/DrawFlowApp/files/program.py")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}
	_, err = cmd.StderrPipe()
	if err != nil {
		panic(err)
	}
	err = cmd.Start()

	response["code"] = strProgram
	response["output"] = copyOutput(stdout)
	cmd.Wait()

	return response, err
}

/* func (ps ProgramService) Update(id string, update_program []byte) (*api.Response, error) {
	txn, _, ctx := ConnectToDatabase()
	q := `query program($id : string) {
			program(func: uid($id)) {
				uid
			}
		}`

	data := map[string]interface{}{}

	json.Unmarshal(update_program, &data)

	data["uid"] = id

	a, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}

	mu := &api.Mutation{
		SetJson: a,
	}

	req := &api.Request{
		Query:     q,
		Vars:      map[string]string{"$id": id},
		Mutations: []*api.Mutation{mu},
		CommitNow: true,
	}

	res, err := txn.Do(ctx, req)
	return res, err
} */

func (ps ProgramService) Show(id string) (map[string]interface{}, error) {
	txn, _, ctx := ConnectToDatabase()

	q := `query program($id : string) {
		program(func: uid($id)) {
		 uid
			program_name
		 nodes
		   {
			 id
			   name
		   data
		   class
		   html
		   typenode
		   inputs {
				   connections{
				   node
			   input
			 }
		   }
		   outputs{
			   connections {
				   node
			 output
		   }
		   }
		   pos_x
		   pos_y
		 }
	   }
	   }`

	data := map[string]interface{}{}

	res, err := txn.QueryWithVars(ctx, q, map[string]string{"$id": id})

	if err != nil {
		fmt.Println(err.Error())
	} else {
		json.Unmarshal(res.GetJson(), &data)
	}

	return data, err
}

func copyOutput(r io.Reader) string {
	scanner := bufio.NewScanner(r)
	response := ""
	for scanner.Scan() {
		response += scanner.Text() + "\n"
	}
	return response
}
