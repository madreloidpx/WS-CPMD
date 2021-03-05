package main

import (
	"fmt"
	"reflect"
)

//Node node object
type Node struct {
	Name     string
	Index    int
	InBinary map[int]bool
	InLink   map[int]bool
	Outlink  map[int]bool
}

type NodeFactory struct {
	Nodes  map[string]Node
	Length int
}

func Nodes() NodeFactory {
	nodes := NodeFactory{
		Nodes:  nil,
		Length: 0,
	}
	return nodes
}

func (nf NodeFactory) CreateNode(name string) Node {
	n := Node{
		Name:     name,
		Index:    nf.Length,
		InBinary: map[int]bool{nf.Length: true},
		InLink:   map[int]bool{nf.Length: true},
		Outlink:  map[int]bool{nf.Length: true},
	}
	nf.Nodes[name] = n
	nf.Length = nf.Length + 1
	return n
}

func (nf NodeFactory) ShowNodes() {
	nodes := reflect.ValueOf(nf.Nodes).MapKeys()
	fmt.Println(nodes)
}

func (n Node) AddInLink(in int) {
	n.InLink[in] = true
}

func (n Node) AddOutLink(out int) {
	n.Outlink[out] = true
}

func (nf NodeFactory) CreateEdge(in, out string) {
	inNode, exist := nf.Nodes[in]
	if exist == false {
		inNode = nf.CreateNode(in)
	}
	outNode, exist := nf.Nodes[out]
	if exist == false {
		outNode = nf.CreateNode(out)
	}
	inNode.AddOutLink(outNode.Index)
	outNode.AddInLink(inNode.Index)
	nf.Nodes[inNode.Name] = inNode
	nf.Nodes[outNode.Name] = outNode
}
