package main

import (
	"testing"
)

func TestAddEdge(t *testing.T) {
	wq := CreateWeakClique()
	n1 := CreateNode(0, "A")
	n2 := CreateNode(1, "B")
	e := CreateEdge(0, "AB", &n2, &n1, 1)
	wq.AddEdge(&e)
	if len(wq.Edges) != 1 {
		t.Error("Edge not added", e, wq)
	}
}

func TestGenerateName(t *testing.T) {
	wq := CreateWeakClique()
	n1 := CreateNode(0, "A")
	n2 := CreateNode(1, "B")
	e1 := CreateEdge(0, "AB", &n2, &n1, 1)
	n3 := CreateNode(2, "C")
	e2 := CreateEdge(2, "BC", &n3, &n2, 1)
	wq.AddEdge(&e1)
	wq.AddEdge(&e2)
	wq.GenerateName()
	if wq.Name != "AB" {
		t.Error("Wrong name generated (1)", wq.Name)
	}
	e3 := CreateEdge(3, "CA", &n1, &n3, 1)
	wq.AddEdge(&e3)
	wq.GenerateName()
	if wq.Name != "ABC" {
		t.Error("Wrong name generated (2)", wq.Name)
	}
}

func TestMarkMerged(t *testing.T) {
	n1 := CreateNode(0, "A")
	n2 := CreateNode(1, "B")
	e1 := CreateEdge(0, "AB", &n2, &n1, 1)
	n3 := CreateNode(2, "C")
	e2 := CreateEdge(1, "BC", &n3, &n2, 1)
	wq1 := CreateWeakClique()
	wq1 = wq1.AddEdge(&e1)
	wq1 = wq1.AddEdge(&e2)
	wq1.MarkMerged()

	if wq1.Merged != true {
		t.Error("Failed to mark clique as merged", wq1.Merged)
	}
}

// func TestCompareWeakCliques(t *testing.T) {
// 	wq1 := CreateWeakClique()
// 	n1 := CreateNode(0, "A")
// 	n2 := CreateNode(1, "B")
// 	e1 := CreateEdge(0, "AB", &n2, &n1, 1)
// 	n3 := CreateNode(2, "C")
// 	e2 := CreateEdge(1, "BC", &n3, &n2, 1)
// 	wq1.AddEdge(&e1)
// 	wq1.AddEdge(&e2)
// 	wq2 := CreateWeakClique()
// 	n4 := CreateNode(3, "D")
// 	n5 := CreateNode(4, "E")
// 	e3 := CreateEdge(2, "DE", &n5, &n4, 1)
// 	wq2.AddEdge(&e2)
// 	wq2.AddEdge(&e3)
// 	finalEdges := []*Edge{&e2}
// 	commonEdges := CompareWeakCliques(wq1, wq2)
// 	for i, edge := range commonEdges {
// 		if edge != finalEdges[i] {
// 			t.Error("Wrong common edges wq1:", wq1, " wq2:", &wq2, "common edges:", commonEdges)
// 		}
// 	}
// }

func TestCheckSameClique(t *testing.T) {
	wq1 := CreateWeakClique()
	n1 := CreateNode(0, "A")
	n2 := CreateNode(1, "B")
	e1 := CreateEdge(0, "AB", &n2, &n1, 1)
	n3 := CreateNode(2, "C")
	e2 := CreateEdge(1, "BC", &n3, &n2, 1)
	wq1 = wq1.AddEdge(&e1)
	wq2 := CreateWeakClique()
	wq2 = wq2.AddEdge(&e1)
	wq2 = wq2.AddEdge(&e2)

	same := CheckSameClique(wq1, wq2)
	if same != true {
		t.Error("Unable to see two same weak cliques, ", wq1, wq2)
	}

	wq3 := CreateWeakClique()
	n4 := CreateNode(3, "D")
	n5 := CreateNode(4, "E")
	e3 := CreateEdge(2, "DE", &n5, &n4, 1)
	wq3 = wq3.AddEdge(&e2)
	wq3 = wq3.AddEdge(&e3)

	same = CheckSameClique(wq1, wq3)
	if same != false {
		t.Error("Unable to see two different weak cliques, ", wq1, wq3)
	}
}

func TestRemoveDuplicates(t *testing.T) {
	var wqList []WeakClique
	wq1 := CreateWeakClique()
	n1 := CreateNode(0, "A")
	n2 := CreateNode(1, "B")
	e1 := CreateEdge(0, "AB", &n2, &n1, 1)
	n3 := CreateNode(2, "C")
	e2 := CreateEdge(1, "BC", &n3, &n2, 1)
	wq1 = wq1.AddEdge(&e1)
	wq1 = wq1.AddEdge(&e2)
	wqList = append(wqList, wq1)
	wq2 := CreateWeakClique()
	wq2 = wq2.AddEdge(&e1)
	wq2 = wq2.AddEdge(&e2)
	wqList = append(wqList, wq2)
	wq3 := CreateWeakClique()
	n4 := CreateNode(3, "D")
	n5 := CreateNode(4, "E")
	e3 := CreateEdge(2, "DE", &n5, &n4, 1)
	wq3 = wq3.AddEdge(&e2)
	wq3 = wq3.AddEdge(&e3)
	wqList = append(wqList, wq3)
	finalList := []WeakClique{wq1, wq3}

	reducedList := RemoveDuplicates(wqList, 0)
	if len(finalList) != len(reducedList) {
		t.Error("Unable to reduce list", finalList, reducedList)
	} else {
		for i, wq := range finalList {
			if CheckSameClique(wq, reducedList[i]) == false {
				t.Error("Unable to reduce list", finalList, reducedList)
			}
		}
	}
}

func TestRemoveMinorCliques(t *testing.T) {
	var wqList []WeakClique
	wq1 := CreateWeakClique()
	n1 := CreateNode(0, "A")
	n2 := CreateNode(1, "B")
	e1 := CreateEdge(0, "AB", &n2, &n1, 1)
	n3 := CreateNode(2, "C")
	e2 := CreateEdge(1, "BC", &n3, &n2, 1)
	wq1 = wq1.AddEdge(&e1)
	wq1 = wq1.AddEdge(&e2)
	wqList = append(wqList, wq1)
	wq2 := CreateWeakClique()
	wq2 = wq2.AddEdge(&e1)
	wq2 = wq2.AddEdge(&e2)
	wq2.MarkMerged()
	wqList = append(wqList, wq2)
	wq3 := CreateWeakClique()
	n4 := CreateNode(3, "D")
	n5 := CreateNode(4, "E")
	e3 := CreateEdge(2, "DE", &n5, &n4, 1)
	wq3 = wq3.AddEdge(&e2)
	wq3 = wq3.AddEdge(&e3)
	wqList = append(wqList, wq3)
	finalList := []WeakClique{wq1, wq3}

	reducedList := RemoveMinorCliques(wqList)
	if len(finalList) != len(reducedList) {
		t.Error("Unable to remove minor cliques", finalList, reducedList)
	} else {
		for i, wq := range finalList {
			if CheckSameClique(wq, reducedList[i]) == false {
				t.Error("Unable to remove minor cliques", finalList, reducedList)
			}
		}
	}
}

func TestListNodes(t *testing.T) {
	wq := CreateWeakClique()
	n1 := CreateNode(0, "A")
	n2 := CreateNode(1, "B")
	e1 := CreateEdge(0, "AB", &n2, &n1, 1)
	n3 := CreateNode(2, "C")
	e2 := CreateEdge(1, "BC", &n3, &n2, 1)
	wq = wq.AddEdge(&e1)
	wq = wq.AddEdge(&e2)
	finalList := []*Node{&n2, &n1, &n3}

	result := wq.ListNodes()
	for i, node := range result {
		if finalList[i] != node {
			t.Error("Unable to get list of nodes in clique", finalList, result)
		}
	}
}
