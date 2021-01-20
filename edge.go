package main

import "fmt"

//Edge object
type Edge struct {
	Index   int
	Name    string
	InNode  *Node
	OutNode *Node
	Weight  float64
}

//CreateEdge if no weight given, use Weight = 1
func CreateEdge(index int, name string, in *Node, out *Node, weight float64) Edge {
	e := Edge{
		Index:   index,
		Name:    name,
		InNode:  in,
		OutNode: out,
		Weight:  weight,
	}
	return e
}

//CheckSimilar check if an edge has same in and out nodes
func (e Edge) CheckSimilar(in string, out string) bool {
	if e.InNode.Name == in && e.OutNode.Name == out {
		return true
	}
	return false
}

//FindEdge finds an edge from an array using CheckSimilar
func FindEdge(edges []*Edge, in string, out string) int {
	for i, edge := range edges {
		if edge.CheckSimilar(in, out) {
			return i
		}
	}
	return -1
}

//MergeEdges merges two edge lists
func MergeEdges(e1, e2 []*Edge) []*Edge {
	var newEdge []*Edge
	for _, edge := range e1 {
		newEdge = append(newEdge, edge)
	}
	for _, edge := range e2 {
		exist := false
		for _, edge2 := range e1 {
			if edge.Name == edge2.Name {
				exist = true
				break
			}
		}
		if exist == false {
			newEdge = append(newEdge, edge)
		}
	}
	return newEdge
}

//CompareEdges returns common edges
func CompareEdges(e1, e2 []*Edge) []*Edge {
	var commonEdges []*Edge
	for _, i := range e1 {
		for _, j := range e2 {
			if i == j {
				commonEdges = append(commonEdges, i)
			}
		}
	}
	return commonEdges
}

//Print print an edge
func (e Edge) Print() string {
	return fmt.Sprintf("Edge(%s%s, %f)", e.OutNode.Name, e.InNode.Name, e.Weight)
}

//PrintList print a list of edges
func PrintList(e []*Edge) string {
	var edgelist string
	for _, edge := range e {
		edgelist = edgelist + " " + edge.Print()
	}
	return fmt.Sprintf("Edges(%s)", edgelist)
}
