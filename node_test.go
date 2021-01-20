package main

import (
	"testing"
)

func TestCreateNode(t *testing.T) {
	var nodeList []Node
	count := len(nodeList)
	name := "A"
	node1 := CreateNode(count, name)
	if node1.Index != 0 && node1.Name != name {
		t.Error("CreateNode (", name, ", ", count, ") not created properly. Node is", node1)
	}
	nodeList = append(nodeList, node1)
	count = len(nodeList)
	name = "B"
	node2 := CreateNode(count, name)
	if node2.Index != 1 && node2.Name != name {
		t.Error("CreateNode (", name, ", ", count, ") not created properly. Node is %s", node2)
	}
	nodeList = append(nodeList, node2)
}

func TestInOutLink(t *testing.T) {
	var nodeList []Node
	// A -> B
	n1 := CreateNode(len(nodeList), "A")
	nodeList = append(nodeList, n1)
	n2 := CreateNode(len(nodeList), "B")
	nodeList = append(nodeList, n2)
	edge := CreateEdge(0, "AB", &n2, &n1, 1)
	n1.InsertOutlink(&edge)
	n2.InsertInlink(&edge)
	if &edge != n1.Outlink[0] {
		t.Error("Edge", edge, "not inserted in outlink of node", n1)
	}
	if &edge != n2.Inlink[0] {
		t.Error("Edge", edge, "not inserted in inlink of node", n2)
	}
}

func TestSearchNode(t *testing.T) {
	var nodeList []*Node
	n1 := CreateNode(len(nodeList), "A")
	nodeList = append(nodeList, &n1)
	n2 := CreateNode(len(nodeList), "B")
	nodeList = append(nodeList, &n2)
	pos := SearchNode(nodeList, "B")
	if pos == -1 {
		t.Error("Node", n2, "was not found")
	}
	if pos != 1 {
		t.Error("Wrong index. Index", pos, "is found when searching", n2)
	}
	pos = SearchNode(nodeList, "C")
	if pos != -1 {
		t.Error("Found a node when it should not exist. Index is", pos, "for node", nodeList[pos])
	}
}

func TestReduceUniqueNodes(t *testing.T) {
	var nodeList []*Node
	n1 := CreateNode(len(nodeList), "A")
	nodeList = append(nodeList, &n1)
	nodeList = append(nodeList, &n1)
	n2 := CreateNode(len(nodeList), "B")
	nodeList = append(nodeList, &n2)
	nodeList = append(nodeList, &n1)
	nodeList = append(nodeList, &n2)
	finalList := []*Node{&n1, &n2}

	result := ReduceUniqueNodes(nodeList, 0)
	for i, node := range result {
		if finalList[i] != node {
			t.Error("Unable to reduce nodelist", finalList, result)
		}
	}
}
