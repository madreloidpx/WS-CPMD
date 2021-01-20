package main

import (
	"fmt"

	"github.com/bradfitz/slice"
)

//WeakClique object
type WeakClique struct {
	Name   string
	Edges  []*Edge
	Nodes  []*Node
	Size   int
	Merged bool
}

//CreateWeakClique creates weak clique
func CreateWeakClique() WeakClique {
	wq := WeakClique{
		Name:   "",
		Size:   0,
		Merged: false,
	}
	return wq
}

//CreateWeakCliqueWithEdge creates weak clique with predefined edges
func CreateWeakCliqueWithEdge(edges []*Edge) WeakClique {
	wq := WeakClique{
		Edges:  edges,
		Size:   len(edges),
		Merged: false,
	}
	wq.GenerateName()
	return wq
}

//AddEdge adds a component edge to weak clique edge array
func (w WeakClique) AddEdge(edge *Edge) WeakClique {
	checkExistEdge := FindEdge(w.Edges, edge.InNode.Name, edge.OutNode.Name)
	if checkExistEdge == -1 {
		w.Edges = append(w.Edges, edge)
		w.Size = w.Size + 1
	}
	return w
}

//MergeWeakCliques merges two weak cliques
func MergeWeakCliques(wq1, wq2 WeakClique) WeakClique {
	newEdges := MergeEdges(wq1.Edges, wq2.Edges)
	newWq := CreateWeakCliqueWithEdge(newEdges)
	newWq.Nodes = ReduceUniqueNodes(append(wq1.Nodes, wq2.Nodes...), 0)
	return newWq
}

//GenerateName updates the weak clique name
func (w *WeakClique) GenerateName() {
	var name string
	for _, edge := range w.Edges {
		name = name + edge.OutNode.Name
	}
	w.Name = name
}

//MarkMerged marks a weak clique as merged
func (w *WeakClique) MarkMerged() {
	w.Merged = true
}

//CompareWeakCliques returns common nodes
func CompareWeakCliques(w1, w2 WeakClique) []*Node {
	return CompareNodes(w1.Nodes, w2.Nodes)
}

//CheckSameClique checks if two cliques are the same clique
func CheckSameClique(w1, w2 WeakClique) bool {
	if len(w1.Nodes) != len(w2.Nodes) {
		return false
	}
	commonNodes := CompareNodes(w1.Nodes, w2.Nodes)
	if len(commonNodes) == len(w1.Nodes) {
		return true
	}
	return false
	// if len(w1.Edges) != len(w2.Edges) {
	// 	return false
	// }
	// for _, e1 := range w1.Edges {
	// 	exist := false
	// 	for _, e2 := range w2.Edges {
	// 		if e1 == e2 {
	// 			exist = true
	// 		}
	// 	}
	// 	if exist == false {
	// 		return false
	// 	}
	// }
	// return true
}

//CheckMinorClique checks if a clique is the minor clique of the other
func CheckMinorClique(w1, w2 WeakClique) bool {
	for _, e1 := range w1.Edges {
		exist := false
		for _, e2 := range w2.Edges {
			if e1 == e2 {
				exist = true
			}
		}
		if exist == false {
			return false
		}
	}
	return true
}

//RemoveDuplicates reduces a weak clique list to unique cliques only; sort by size beforehand first before using
func RemoveDuplicates(c []WeakClique, start int) []WeakClique {
	if len(c) == 0 {
		return nil
	}
	// if start == 0 {
	// 	fmt.Println("Sorting...")
	// 	slice.Sort(c, func(i, j int) bool {
	// 		return c[i].Size > c[j].Size
	// 	})
	// }
	// fmt.Println("Reducing duplicates; len:", len(c), "start:", start)
	if len(c) == start {
		return c
	}
	unique := c[:start+1]
	// fmt.Println(unique)
	for i := start + 1; i < len(c); i++ {
		if CheckSameClique(c[i], c[start]) == false {
			// fmt.Println(c[i].Print())
			unique = append(unique, c[i])
		}
	}
	return RemoveDuplicates(unique, start+1)
}

//RemoveDuplicatesForMinorClique reduces a weak clique list to unique cliques only; sort by size beforehand first before using
func RemoveDuplicatesForMinorClique(c []WeakClique, start int) []WeakClique {
	if len(c) == 0 {
		return nil
	}
	if start == 0 {
		fmt.Println("Sorting...")
		slice.Sort(c, func(i, j int) bool {
			return c[i].Size > c[j].Size
		})
	}
	// fmt.Println("Reducing duplicates; len:", len(c), "start:", start)
	if len(c) == start {
		return c
	}
	unique := c[:start+1]
	// fmt.Println(unique)
	for i := start + 1; i < len(c); i++ {
		if CheckMinorClique(c[i], c[start]) == false {
			// fmt.Println(c[i].Print())
			unique = append(unique, c[i])
		}
	}
	return RemoveDuplicatesForMinorClique(unique, start+1)
}

//RemoveMinorCliques removes cliques that are merged inside bigger cliques
func RemoveMinorCliques(c []WeakClique) []WeakClique {
	if len(c) == 0 {
		return nil
	}
	fmt.Println("Removing minor cliques inside bigger cliques; size:", len(c))
	var newList []WeakClique
	for _, wq := range c {
		if wq.Merged == false {
			newList = append(newList, wq)
		}
	}
	// fmt.Println("After removing minor cliques,", len(newList))
	return newList
}

//Print print a weak clique
func (w WeakClique) Print() string {
	return fmt.Sprintf("WQ(%s)", w.Name)
}

//PrintWQList prints a list of weak cliques
func PrintWQList(wqs []WeakClique) {
	fmt.Println("Cliques:")
	for _, wq := range wqs {
		fmt.Printf("WeakClique(" + wq.Name + ")\n")
	}
}

//ListNodes lists the nodes within a weak clique
func (w WeakClique) ListNodes() []*Node {
	var nodelist []*Node
	for _, edge := range w.Edges {
		nodelist = append(nodelist, edge.InNode)
		nodelist = append(nodelist, edge.OutNode)
	}
	uniqueNodes := ReduceUniqueNodes(nodelist, 0)
	return uniqueNodes
}
