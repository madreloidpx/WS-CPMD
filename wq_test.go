package main

import (
	"testing"
)

func TestCreateWeakCliqueFactory(t *testing.T) {
	weakCliques := WeakCliques()
	if weakCliques.Length != 0 {
		t.Error("Weak Clique Factory not created")
	}
}

func TestCreateWeakClique(t *testing.T) {
	nodes := Nodes()
	nodes.CreateEdge("a", "b")
	nodes.CreateEdge("a", "c")
	nodes.CreateEdge("c", "b")
	nodes.CreateEdge("a", "d")
	var nodelist []Node
	for i, node := range nodes.Nodes {
		if i != "d" {
			nodelist = append(nodelist, node)
		}
	}
	weakCliques := WeakCliques()
	weakCliques.CreateWeakClique(nodelist)
	_, exist := weakCliques.WeakCliques["7"]
	if exist == false {
		t.Error("Weak Clique not made.", weakCliques)
	}
}

func TestMergeCliques(t *testing.T) {
	nodes := Nodes()
	nodes.CreateEdge("a", "b")
	nodes.CreateEdge("a", "c")
	nodes.CreateEdge("c", "b")
	nodes.CreateEdge("a", "d")
	list1 := []Node{nodes.Nodes["a"], nodes.Nodes["b"]}
	list2 := []Node{nodes.Nodes["c"], nodes.Nodes["d"]}
	weakcliques := WeakCliques()
	weakcliques.CreateWeakClique(list1)
	weakcliques.CreateWeakClique(list2)
	wq1 := weakcliques.WeakCliques["3"]
	wq2 := weakcliques.WeakCliques["C"]
	newWQL := WeakCliques()
	newWQL.MergeCliques(wq1, wq2)
	_, exist := newWQL.WeakCliques["F"]
	if exist == false {
		t.Error("Weak Cliques didn't merge.", wq1, wq2, newWQL)
	}
}

func TestHexName(t *testing.T) {
	m1 := []int{0, 1, 2, 3}
	name := HexName(m1)
	if name != "F" {
		t.Error("Hex not converted properly", name, m1)
	}
	m2 := []int{4, 6, 5, 2, 7}
	name = HexName(m2)
	if name != "F4" {
		t.Error("Hex not converted properly", name, m2)
	}
	m3 := []int{20, 0}
	name = HexName(m3)
	if name != "100001" {
		t.Error("Hex not converted properly", name, m3)
	}
}
