package main

import (
	"testing"
)

func TestCreateEdge(t *testing.T) {
	var edgeList []Edge
	// A -> B
	n1 := CreateNode(0, "A")
	n2 := CreateNode(1, "B")
	e1 := CreateEdge(len(edgeList), "AB", &n2, &n1, 1)
	if e1.Index != 0 && e1.InNode != &n2 && e1.OutNode != &n1 {
		t.Error("Edge", &e1, "not saved properly")
	}
}

func TestCheckSimilar(t *testing.T) {
	n1 := CreateNode(0, "A")
	n2 := CreateNode(1, "B")
	e1 := CreateEdge(0, "AB", &n2, &n1, 1)
	similar := e1.CheckSimilar("B", "A")
	if similar == false {
		t.Error("Edge", &e1, "not similar to itself")
	}
}

func TestFindEdge(t *testing.T) {
	var edgeList []*Edge
	n1 := CreateNode(0, "A")
	n2 := CreateNode(1, "B")
	n3 := CreateNode(2, "C")
	e1 := CreateEdge(len(edgeList), "AB", &n2, &n1, 1)
	edgeList = append(edgeList, &e1)
	e2 := CreateEdge(len(edgeList), "BC", &n3, &n2, 1)
	edgeList = append(edgeList, &e2)
	pos := FindEdge(edgeList, "C", "B")
	if pos == -1 {
		t.Error("Edge", &e2, "not found")
	}
	if pos != 1 {
		t.Error("Wrong index (", pos, ") found for Edge", &e2)
	}
	pos = FindEdge(edgeList, "D", "E")
	if pos != -1 {
		t.Error("Edge found that should not exist for pos", pos, "as Edge", &edgeList[pos])
	}
}

func TestMergeEdges(t *testing.T) {
	var el1 []*Edge
	var el2 []*Edge
	n1 := CreateNode(0, "A")
	n2 := CreateNode(1, "B")
	n3 := CreateNode(2, "C")
	e1 := CreateEdge(len(el1), "AB", &n2, &n1, 1)
	el1 = append(el1, &e1)
	e2 := CreateEdge(len(el1), "BC", &n3, &n2, 1)
	el1 = append(el1, &e2)
	e3 := CreateEdge(len(el2), "BC", &n3, &n2, 1)
	el2 = append(el2, &e3)
	e4 := CreateEdge(len(el2), "CA", &n1, &n3, 1)
	el2 = append(el2, &e4)
	finalEdges := []*Edge{&e1, &e2, &e4}
	mergedEdges := MergeEdges(el1, el2)
	for i, edge := range mergedEdges {
		if edge != finalEdges[i] {
			t.Error("Edges not merged e1:", &el1, " e2:", &el2, "merged edges:", mergedEdges)
		}
	}

}

func TestCompareEdges(t *testing.T) {
	var el1 []*Edge
	var el2 []*Edge
	n1 := CreateNode(0, "A")
	n2 := CreateNode(1, "B")
	n3 := CreateNode(2, "C")
	e1 := CreateEdge(len(el1), "AB", &n2, &n1, 1)
	el1 = append(el1, &e1)
	e2 := CreateEdge(len(el1), "BC", &n3, &n2, 1)
	el1 = append(el1, &e2)
	e3 := CreateEdge(len(el2), "BC", &n3, &n2, 1)
	el2 = append(el2, &e3)
	e4 := CreateEdge(len(el2), "CA", &n1, &n3, 1)
	el2 = append(el2, &e4)
	finalEdges := []*Edge{&e2}
	commonEdges := CompareEdges(el1, el2)
	for i, edge := range commonEdges {
		if edge != finalEdges[i] {
			t.Error("Wrong common edges e1:", &el1, " e2:", &el2, "common edges:", commonEdges)
		}
	}

}
