package models

import (
	"fmt"
	"strconv"
	"strings"
)

type Program struct {
	Uid   string `json:"uid,omitempty"`
	Name  string `json:"name,omitempty"`
	Nodes []Node `json:"nodes,omitempty"`
}

func (p Program) WriteProgram() string {
	block := p.Nodes[0].Outputs[0].Connections

	a := ""
	for _, item := range block {
		Id, err := strconv.Atoi(item.Node)
		if err != nil {
			fmt.Errorf(err.Error())
		}

		aux_node := p.FindNode(Id)
		a += p.Compare(aux_node)
	}
	return a
}

func (p Program) Compare(node Node) string {
	a := ""

	switch node.Name {
	case "for":
		a += p.ProgramLoopFor(node)
	case "conditional":
		a += p.ProgramConditional(node)
	case "assign":
		a += p.ProgramAssignements(node) + "\n"
	case "operator":
		a += p.ProgramAssignements(node) + "\n"

	}
	return a
}

func (p Program) ProgramLoopFor(node Node) string {
	a := "for i in range(" + node.Data["begin"] + "," + node.Data["end"] + "):\n"

	for _, item := range node.Outputs[0].Connections {
		Id, err := strconv.Atoi(item.Node)
		if err != nil {
			fmt.Errorf(err.Error())
		}
		aux_node := p.FindNode(Id)
		val := p.Compare(aux_node)
		a += val
	}
	return strings.ReplaceAll(a, "\n", "\n\t")
}

func (p Program) ProgramConditional(node Node) string {
	a := "if " + node.Data["variable"] + node.Data["type"] + node.Data["constant"] + ":\n"
	Id, err := strconv.Atoi(node.Outputs[0].Connections[0].Node)
	if err != nil {
		fmt.Errorf(err.Error())
	}
	if_st := p.FindNode(Id)
	a += p.ProgramIf(if_st) + "else:\n"
	Id, err = strconv.Atoi(node.Outputs[1].Connections[0].Node)
	if err != nil {
		fmt.Errorf(err.Error())
	}
	else_st := p.FindNode(Id)
	a += p.ProgramElse(else_st) + "\n"

	return a
}

func (p Program) ProgramIf(node Node) string {
	a := ""

	for _, item := range node.Outputs[0].Connections {
		Id, err := strconv.Atoi(item.Node)
		if err != nil {
			fmt.Errorf(err.Error())
		}
		AuxNode := p.FindNode(Id)
		a += "\t" + p.Compare(AuxNode) + "\n"
	}
	return a
}

func (p Program) ProgramElse(node Node) string {
	a := ""

	for _, item := range node.Outputs[0].Connections {
		Id, err := strconv.Atoi(item.Node)
		if err != nil {
			fmt.Errorf(err.Error())
		}
		AuxNode := p.FindNode(Id)
		a += "\t" + p.Compare(AuxNode) + "\n"
	}

	return a
}
func (p Program) ProgramAssignements(node Node) string {
	a := ""

	if len(node.Outputs) == 0 {
		a += p.EvaluateProgram(node)
	} else {
		Id, err := strconv.Atoi(node.Outputs[0].Connections[0].Node)
		if err != nil {
			fmt.Errorf(err.Error())
		}
		a += p.ProgramAssignements(p.FindNode(Id))
		a += p.EvaluateProgram(node)
		Id, err = strconv.Atoi(node.Outputs[1].Connections[0].Node)
		if err != nil {
			fmt.Errorf(err.Error())
		}
		a += p.ProgramAssignements(p.FindNode(Id))
	}

	return a
}

func (p Program) EvaluateProgram(node Node) string {
	a := ""
	switch node.Name {
	case "variable":
		a = node.Data["name"]
	case "constant":
		a = node.Data["value"]
	case "assign":
		a = "="
	case "operator":
		a = node.Data["operator"]
	}
	return a
}

func (p Program) FindNode(id int) Node {
	for _, item := range p.Nodes {
		if id == item.Id {
			return item
		}
	}
	return Node{}
}

type Node struct {
	Id       int               `json:"id,omitempty"`
	Name     string            `json:"name,omitempty"`
	Data     map[string]string `json:"data,omitempty"`
	Class    string            `json:"class,omitempty"`
	Html     string            `json:"html,omitempty"`
	TypeNode string            `json:"typenode,omitempty"`
	Inputs   []Input           `json:"inputs,omitempty"`
	Outputs  []Output          `json:"outputs,omitempty"`
	PosX     int               `json:"pos_x,omitempty"`
	PosY     int               `json:"pos_y,omitempty"`
}

type Input struct {
	Connections []InputConn `json:"connections,omitempty"`
}

type InputConn struct {
	Node  string `json:"node,omitempty"`
	Input string `json:"input,omitempty"`
}

type Output struct {
	Connections []OutputConn `json:"connections,omitempty"`
}

type OutputConn struct {
	Node   string `json:"node,omitempty"`
	Output string `json:"output,omitempty"`
}
