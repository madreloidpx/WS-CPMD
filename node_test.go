package main

import (
	"fmt"
	"testing"
)

func TestCreateNodeFactory(t *testing.T) {
	nodes := Nodes()
	if nodes.Length != 0 {
		t.Error("Node Factory not created.", nodes.Nodes)
	}
}

func TestCreateNode(t *testing.T) {
	nodes := Nodes()
	nodes.CreateNode("a")
	node, exist := nodes.Nodes["a"]
	if exist == false || node.Name != "a" || node.Index != 0 {
		t.Error("Node not created.")
		node.PrintNode()
	}
	nodes.CreateNode("b")
	node, exist = nodes.Nodes["b"]
	if exist == false || node.Name != "b" || node.Index != 1 {
		t.Error("Node not created.")
		node.PrintNode()
	}
}

func TestAddWeakClique(t *testing.T) {
	nodes := Nodes()
	nodes.CreateEdge("a", "b")
	nodes.CreateEdge("a", "c")
	nodes.CreateEdge("c", "b")
	var nodelist []*Node
	for _, node := range nodes.Nodes {
		nodelist = append(nodelist, &node)
	}
	node := nodes.IndexMap[0]
	weakClique := CreateWeakClique(nodelist)
	nodeMembers := make([]int, 0, len(nodelist))
	for _, n := range nodelist {
		nodeMembers = append(nodeMembers, n.Index)
	}
	hexname := HexName(nodeMembers)
	node.AddWeakClique(hexname, &weakClique)
	_, exist := node.WeakCliques[hexname]
	if exist == false {
		t.Error("WeakClique not added.")
		node.PrintNode()
	}
	if weakClique.Parent != node {
		t.Error("Parent not added to weak clique.", weakClique)
	}
}

func TestAddWQMark(t *testing.T) {
	nodes := Nodes()
	nodes.CreateEdge("a", "b")
	nodes.CreateEdge("a", "c")
	nodes.CreateEdge("c", "b")
	var nodelist []*Node
	for _, node := range nodes.Nodes {
		nodelist = append(nodelist, &node)
	}
	node := nodes.IndexMap[0]
	weakClique := CreateWeakClique(nodelist)
	nodeMembers := make([]int, 0, len(nodelist))
	for _, n := range nodelist {
		nodeMembers = append(nodeMembers, n.Index)
	}
	hexname := HexName(nodeMembers)
	node.AddWeakClique(hexname, &weakClique)
	node.AddWQMark()
	if node.Complete == false {
		t.Error("Node didn't mark complete.")
		node.PrintNode()
	}
}

func TestCreateEdge(t *testing.T) {
	nodes := Nodes()
	nodes.CreateEdge("a", "b")
	node, exist := nodes.Nodes["a"]
	if exist == false || node.Name != "a" || node.Index != 0 || node.OutLink[1] == false {
		t.Error("Edge not created.")
		node.PrintNode()
		fmt.Println("\nNodes:", nodes.ShowNodes())
	}
	node, exist = nodes.Nodes["b"]
	if exist == false || node.Name != "b" || node.Index != 1 || node.InLink[0] == false {
		t.Error("Edge not created.")
		node.PrintNode()
		fmt.Println("\nNodes:", nodes.ShowNodes())
	}

	nodes.CreateEdge("b", "c")
	node, exist = nodes.Nodes["c"]
	if exist == false || node.Name != "c" || node.Index != 2 || node.InLink[1] == false {
		t.Error("Edge not created.")
		node.PrintNode()
		fmt.Println("\nNodes:", nodes.ShowNodes())
	}
	node, exist = nodes.Nodes["b"]
	if exist == false || node.Name != "b" || node.Index != 1 || node.OutLink[2] == false {
		t.Error("Edge not updated.")
		node.PrintNode()
		fmt.Println("\nNodes:", nodes.ShowNodes())
	}

	nodes.CreateEdge("c", "a")
	node, exist = nodes.Nodes["c"]
	if exist == false || node.Name != "c" || node.Index != 2 || node.OutLink[0] == false {
		t.Error("Edge not updated.")
		node.PrintNode()
		fmt.Println("\nNodes:", nodes.ShowNodes())
	}
	node, exist = nodes.Nodes["a"]
	if exist == false || node.Name != "a" || node.Index != 0 || node.InLink[2] == false {
		t.Error("Edge not updated.")
		node.PrintNode()
		fmt.Println("\nNodes:", nodes.ShowNodes())
	}
}
