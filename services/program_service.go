package services

import (
	"DrawFlowApp/dao"
	"DrawFlowApp/models"
	"os/exec"

	"bufio"
	"context"
	"encoding/json"
	"io"
	"log"

	"github.com/dgraph-io/dgo/v2"
	"github.com/dgraph-io/dgo/v2/protos/api"
	"google.golang.org/grpc"
)

type ProgramService struct {
	URL string
}

func (ps ProgramService) ConnectToDatabase() (*dgo.Txn, *grpc.ClientConn, context.Context) {
	ctx := context.TODO()

	/* Conexion a la base de datos "localhost:9080" */
	conn, err := grpc.Dial(ps.URL, grpc.WithInsecure())
	if err != nil {
		log.Fatal("failed to dial ", err)
	}
	defer conn.Close()
	dgraphClient := dgo.NewDgraphClient(api.NewDgraphClient(conn))

	txn := dgraphClient.NewTxn()
	defer txn.Commit(ctx)

	return txn, conn, ctx
}

func (ps ProgramService) Store(new_program []byte) error {
	txn, conn, ctx := ps.ConnectToDatabase()

	mu := &api.Mutation{
		SetJson: new_program,
	}

	_, err := txn.Mutate(ctx, mu)

	conn.Close()

	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
func (ps ProgramService) Execute(execute_program []byte) string {
	var program models.Program
	path := "./files/program.py"

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
	if err != nil {
		panic(err)
	}

	response := copyOutput(stdout)
	cmd.Wait()

	return response

}
func (ps ProgramService) Index() {

}

func (ps ProgramService) Update(id string, update_program models.Program) {

}

func (ps ProgramService) Show(id string) {

}

func copyOutput(r io.Reader) string {
	scanner := bufio.NewScanner(r)
	response := ""
	for scanner.Scan() {
		response += scanner.Text() + "\n"
	}
	return response
}
