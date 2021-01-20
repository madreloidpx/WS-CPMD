package main

import (
	"testing"
)

func TestReadNodeFile(t *testing.T) {
	justCheckValue := true
	data := ReadNodeFile("testGraph.dat")
	if justCheckValue == true {
		t.Error("Graph ReadNodeFile", data)
	}
}

func TestCreateGraphByFile(t *testing.T) {
	justCheckValue := true
	graph := CreateGraphByFile("testGraph.dat")
	if justCheckValue == true {
		graph.Print()
		t.Error("Graph CreateGraphFile")
	}
}

func TestCreateWeakCliques(t *testing.T) {
	justCheckValue := true
	graph := CreateGraphByFile("testGraph.dat")
	graph.CreateWeakCliques()
	if justCheckValue == true {
		graph.Print()
		t.Error("Graph CreateWeakCliques")
	}
}

func TestMergeCliques(t *testing.T) {
	justCheckValue := true
	graph := CreateGraphByFile("testGraph.dat")
	graph.CreateWeakCliques()
	graph.MergeCliques(1)
	if justCheckValue == true {
		graph.Print()
		t.Error("Graph MergeCliques")
	}
}

func TestFormCommunity(t *testing.T) {
	justCheckValue := true
	graph := CreateGraphByFile("testGraph1.dat")
	graph.CreateWeakCliques()
	graph.FormCommunity()
	if justCheckValue == true {
		graph.Print()
		t.Error("Graph FormCommunity")
	}
}

func TestGenerateCommunityFile(t *testing.T) {
	justCheckValue := true
	graph := CreateGraphByFile("tests/10/community4.dat")
	graph.CreateWeakCliques()
	graph.FormCommunity()
	if justCheckValue == true {
		graph.Print()
		graph.GenerateCommunityFile("results/testResult.dat")
		t.Error("Graph GenerateCommunityFile")
	}
}
