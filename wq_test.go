package main

import (
	"fmt"
	"testing"
)

func TestCreateWeakClique(t *testing.T) {
	nodes := Nodes()
	nodes.CreateEdge("a", "b")
	nodes.CreateEdge("a", "c")
	nodes.CreateEdge("c", "b")
	nodelist := make([]*Node, 0)
	for _, node := range nodes.IndexMap {
		nodelist = append(nodelist, node)
	}
	// weakCliques := WeakCliques()
	fmt.Println(nodelist)
	weakClique := CreateWeakClique(nodelist)
	nodeMembers := make([]int, 0)
	for _, n := range nodelist {
		nodeMembers = append(nodeMembers, n.Index)
	}
	hexname := HexName(nodeMembers)
	if hexname != "7" {
		t.Error("Weak clique generated incorrect.", nodes, hexname, weakClique)
	}
}

func TestMergeCliques(t *testing.T) {
	nodes := Nodes()
	nodes.CreateEdge("a", "b")
	nodes.CreateEdge("a", "c")
	nodes.CreateEdge("c", "b")
	nodes.CreateEdge("a", "d")
	a, _ := nodes.Nodes["a"]
	b, _ := nodes.Nodes["b"]
	c, _ := nodes.Nodes["c"]
	d, _ := nodes.Nodes["d"]
	list1 := []*Node{&a, &b}
	list2 := []*Node{&c, &d}
	wq1 := CreateWeakClique(list1)
	wq2 := CreateWeakClique(list2)
	wqMerged := MergeCliques(&wq1, &wq2)
	nodeMembers := make([]int, 0, len(wqMerged.Members))
	for index, _ := range wqMerged.Members {
		nodeMembers = append(nodeMembers, index)
	}
	hexname := HexName(nodeMembers)
	if hexname != "F" {
		t.Error("Weak cliques not merged properly", hexname, wq1, wq2, wqMerged)
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

func TestMarkMerged(t *testing.T) {
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
	weakClique.MarkMerged()
	if node.Complete == false {
		t.Error("Node didn't mark complete.")
		node.PrintNode()
	}
}

func TestMergeWeakCliqueList(t *testing.T) {
	nodes := Nodes()
	nodes.CreateEdge("a", "b")
	nodes.CreateEdge("a", "c")
	nodes.CreateEdge("c", "b")
	nodes.CreateEdge("d", "a")
	nodes.CreateEdge("b", "d")
	a, _ := nodes.Nodes["a"]
	b, _ := nodes.Nodes["b"]
	c, _ := nodes.Nodes["c"]
	d, _ := nodes.Nodes["d"]
	list1 := []*Node{&a, &b, &c, &d}
	list2 := []*Node{&a, &b, &c}
	wq1 := CreateWeakClique(list1)
	wq1.Parent = &c
	wq2 := CreateWeakClique(list2)
	wq2.Parent = &a
	wq3 := CreateWeakClique(list2)
	wq3.Parent = &a
	wqlist1 := make(map[string]*WeakClique)
	wqlist1["F"] = &wq1
	wqlist2 := make(map[string]*WeakClique)
	wqlist2["7"] = &wq2
	wqlist2["F"] = &wq3
	merged := MergeWeakCliqueList(wqlist1, wqlist2)
	_, exist := merged["F"]
	_, exist2 := merged["7"]
	if exist == false || exist2 == false {
		t.Error("merge lists didn't merge properly", merged)
	}
	if wq3.Merged == false {
		t.Error("merge didn't mark merge duplicate", wq3)
	}
	if a.Complete == true || c.Complete == true {
		t.Error("merge incorrectly marked parent as complete", a, wq3)
	}
}

func TestCheckMerge(t *testing.T) {
	nodes := Nodes()
	nodes.CreateEdge("a", "b")
	nodes.CreateEdge("a", "c")
	nodes.CreateEdge("c", "b")
	nodes.CreateEdge("d", "a")
	nodes.CreateEdge("b", "d")
	a, _ := nodes.Nodes["a"]
	b, _ := nodes.Nodes["b"]
	c, _ := nodes.Nodes["c"]
	d, _ := nodes.Nodes["d"]
	list1 := []*Node{&a, &b, &c, &d}
	list2 := []*Node{&a, &b, &c}
	wq1 := CreateWeakClique(list1)
	wq2 := CreateWeakClique(list2)
	checkMerge := CheckMerge(&wq1, &wq2)
	if checkMerge == false {
		t.Error("Check merge incorrect", checkMerge)
	}
}

func TestMergeCliquesIteration(t *testing.T) {
	nodes := Nodes()
	nodes.CreateEdge("a", "b")
	nodes.CreateEdge("a", "c")
	nodes.CreateEdge("c", "b")
	nodes.CreateEdge("d", "a")
	nodes.CreateEdge("b", "d")
	a, _ := nodes.Nodes["a"]
	b, _ := nodes.Nodes["b"]
	c, _ := nodes.Nodes["c"]
	d, _ := nodes.Nodes["d"]
	list1 := []*Node{&a, &b, &c, &d}
	list2 := []*Node{&a, &b, &c}
	wq1 := CreateWeakClique(list1)
	c.AddWeakClique("F", &wq1)
	wq2 := CreateWeakClique(list2)
	a.AddWeakClique("7", &wq2)
	wq3 := CreateWeakClique(list2)
	a.AddWeakClique("F", &wq3)
	wqlist1 := make(map[string]*WeakClique)
	wqlist1["F"] = &wq1
	wqlist2 := make(map[string]*WeakClique)
	wqlist2["7"] = &wq2
	wqlist2["F"] = &wq3
	mergedList := MergeWeakCliqueList(wqlist1, wqlist2)
	MergeCliquesIteration(mergedList)
	if a.Complete == false || c.Complete == true {
		t.Error("Incorrectly marked wq parents", a, c)
	}
	if wq1.Merged == true || wq2.Merged == false || wq3.Merged == false {
		t.Error("Incorrect marking of wq.", wq1, wq2, wq3)
	}
}
