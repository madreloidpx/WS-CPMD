package main

import (
	"fmt"
)

//Node node object
type Node struct {
	Name    string
	Index   int
	InLink  map[int]bool
	OutLink map[int]bool
}

type NodeFactory struct {
	Nodes    map[string]Node
	IndexMap map[int]Node
	Length   int
}

func Nodes() NodeFactory {
	nodes := NodeFactory{
		Nodes:    make(map[string]Node),
		IndexMap: make(map[int]Node),
		Length:   0,
	}
	return nodes
}

func (nf *NodeFactory) CreateNode(name string) Node {
	n := Node{
		Name:    name,
		Index:   nf.Length,
		InLink:  map[int]bool{nf.Length: true},
		OutLink: map[int]bool{nf.Length: true},
	}
	nf.Nodes[name] = n
	nf.IndexMap[nf.Length] = n
	nf.Length = nf.Length + 1
	return n
}

func (nf NodeFactory) ShowNodes() []string {
	nodes := make([]string, 0, nf.Length)
	for k := range nf.Nodes {
		nodes = append(nodes, k)
	}
	return nodes
}

func (n *Node) AddInLink(in int) {
	n.InLink[in] = true
}

func (n *Node) AddOutLink(out int) {
	n.OutLink[out] = true
}

func (n Node) ShowInLink() []int {
	links := make([]int, 0, len(n.InLink))
	for k := range n.InLink {
		links = append(links, k)
	}
	return links
}

func (n Node) ShowOutLink() []int {
	links := make([]int, 0, len(n.OutLink))
	for k := range n.OutLink {
		links = append(links, k)
	}
	return links
}

func (n Node) PrintNode() {
	fmt.Print("Node:", n.Name)
	fmt.Print(" Index:", n.Index)
	fmt.Print(" In:", n.ShowInLink())
	fmt.Print(" Out:", n.ShowOutLink())
}

func (nf *NodeFactory) CreateEdge(in, out string) {
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

func GenerateNodeData(data [][]string) NodeFactory {
	nodes := Nodes()
	for _, nodepair := range data {
		nodes.CreateEdge(nodepair[0], nodepair[1])
	}
	return nodes
}
