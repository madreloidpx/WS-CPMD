package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/bradfitz/slice"
)

//Graph object
type Graph struct {
	Nodes       []*Node
	Edges       []*Edge
	WeakCliques []WeakClique
	Communities []WeakClique
}

type nodeCom struct {
	node      *Node
	community []int
}

//ReadNodeFile reads the node file and returns the data in a string array
func ReadNodeFile(filename string) [][]string {
	file, err := os.Open(filename)
	Handle(err)
	defer file.Close()

	var data [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		data = append(data, fields)
	}
	err = scanner.Err()
	Handle(err)

	return data
}

//CreateGraphByFile creates graph with file in forman (out in weight)
func CreateGraphByFile(filename string) Graph {
	data := ReadNodeFile(filename)
	var nodeList []*Node
	var edgeList []*Edge
	for _, line := range data {
		out := line[0]
		in := line[1]
		var weight float64
		var err error
		if len(line) == 3 {
			weight, err = strconv.ParseFloat(line[2], 64)
			Handle(err)
		} else {
			weight = 1
		}
		outIndex := SearchNode(nodeList, out)
		inIndex := SearchNode(nodeList, in)
		var outNode, inNode *Node
		if outIndex == -1 {
			newNode := CreateNode(len(nodeList), out)
			nodeList = append(nodeList, &newNode)
			outNode = nodeList[len(nodeList)-1]
		} else {
			outNode = nodeList[outIndex]
		}
		if inIndex == -1 {
			newNode := CreateNode(len(nodeList), in)
			nodeList = append(nodeList, &newNode)
			inNode = nodeList[len(nodeList)-1]
		} else {
			inNode = nodeList[inIndex]
		}
		edge := CreateEdge(len(edgeList), out+in, inNode, outNode, weight)
		edgeList = append(edgeList, &edge)
		inNode.InsertInlink(&edge)
		outNode.InsertOutlink(&edge)
		// fmt.Println(inNode.Print(), outNode.Print())
	}
	graph := Graph{
		Nodes: nodeList,
		Edges: edgeList,
	}
	return graph
}

//CreateWeakCliques generates weak cliques from graph's edge and node data
func (g *Graph) CreateWeakCliques() {
	fmt.Println("Creating WQ")
	for _, node := range g.Nodes {
		// fmt.Println("In Node:", node.Print())
		for _, out := range node.Outlink {
			wq := WeakClique{Size: 0}
			// fmt.Println("In Out Edge", out.Print())
			inNodeOfOut := out.InNode
			// fmt.Println("inNodeOfOut:", inNodeOfOut.Print())
			for _, in := range node.Inlink {
				// fmt.Println("In In Edge:", in.Print())
				outNodeOfIn := in.OutNode
				// fmt.Println("outNodeOfIn:", outNodeOfIn.Print())
				// fmt.Println("Finding edge: ", PrintList(inNodeOfOut.Outlink), outNodeOfIn.Name, inNodeOfOut.Name)
				checkEdge := FindEdge(inNodeOfOut.Outlink, outNodeOfIn.Name, inNodeOfOut.Name)
				// fmt.Println("checkEdge:", checkEdge)
				if checkEdge > -1 {
					// fmt.Println("Adding:", out.Print(), in.Print(), inNodeOfOut.Outlink[checkEdge].Print())
					wq = wq.AddEdge(out)
					wq = wq.AddEdge(in)
					wq = wq.AddEdge(inNodeOfOut.Outlink[checkEdge])
				}
			}
			if wq.Size > 0 {
				wq.GenerateName()
				wq.Nodes = wq.ListNodes()
				// fmt.Println("Created wq", wq.Name)
				// fmt.Println("Appended wq:", wq.Print())
				g.WeakCliques = append(g.WeakCliques, wq)
			}
		}
	}
	g.WeakCliques = RemoveDuplicates(g.WeakCliques, 0)
}

//FormCommunity generate communities by merging the weak cliques in increasing k
func (g *Graph) FormCommunity() {
	fmt.Println("Forming Communities", len(g.WeakCliques))
	k := 2
	for k < len(g.Nodes) {
		if len(g.WeakCliques) == 0 {
			break
		}
		fmt.Println("k: ", k)
		g.MergeCliques(k)
		g.WeakCliques = RemoveDuplicates(g.WeakCliques, 0)
		k = k + 1
	}
	g.Communities = RemoveDuplicatesForMinorClique(g.Communities, 0)
}

// MergeCliques generate bigger cliques from graph's weak clique data
func (g *Graph) MergeCliques(k int) {
	fmt.Println("Merging Cliques")
	fmt.Println("Current WQs: ", len(g.WeakCliques))
	var wqList []WeakClique
	for i := 0; i < len(g.WeakCliques); i++ {
		for j := i + 1; j < len(g.WeakCliques); j++ {
			// fmt.Println("Comparing WQs:", g.WeakCliques[i].Print(), g.WeakCliques[j].Print())
			commonNodes := CompareWeakCliques(g.WeakCliques[i], g.WeakCliques[j])
			if len(commonNodes) == len(g.WeakCliques[i].Nodes) {
				continue
			}
			if len(commonNodes) >= k {
				wqNew := MergeWeakCliques(g.WeakCliques[i], g.WeakCliques[j])
				// fmt.Println("Merging WQ =>", wqNew.Print())
				g.WeakCliques[i].MarkMerged()
				g.WeakCliques[j].MarkMerged()
				wqList = append(wqList, wqNew)
			}
		}
		if g.WeakCliques[i].Merged == false {
			// fmt.Println("Adding to community:", g.WeakCliques[i].Print())
			g.Communities = append(g.Communities, g.WeakCliques[i])
		}
	}
	g.WeakCliques = wqList
	fmt.Println("After Merging: WQs =", len(g.WeakCliques), "Coms =", len(g.Communities))
}

//Print print the entire graph details
func (g Graph) Print() {
	fmt.Println("Nodes: ")
	for _, n := range g.Nodes {
		fmt.Println(n.Print())
	}
	fmt.Println("Edges: ")
	for _, e := range g.Edges {
		fmt.Println(e.Print())
	}
	fmt.Printf("Nodes Size: %d\n", len(g.Nodes))
	fmt.Printf("Edges Size: %d\n", len(g.Edges))
	fmt.Println("Community/Cliques: ")
	PrintWQList(g.Communities)
	fmt.Println("Community Size: ", len(g.Communities))
}

//GenerateCommunityFile creates the community.dat file for the given network data
func (g Graph) GenerateCommunityFile(filename string) {
	fmt.Println("Generating Community File")
	nodeComList := g.createNodeCommunityList()
	nodeComList = sortNodeComList(nodeComList)
	printCommunityGivenNodeComData(filename, nodeComList)
}

func sortNodeComList(nodeComList []nodeCom) []nodeCom {
	newList := nodeComList
	slice.Sort(newList, func(i, j int) bool {
		// if _, err := strconv.Atoi(newList[i].node.Name); err == nil {
		// 	i, _ := strconv.Atoi(newList[i].node.Name)
		// 	j, _ := strconv.Atoi(newList[i].node.Name)
		// 	return i < j
		// }
		return newList[i].node.Name < newList[j].node.Name
	})
	return newList
}

func printCommunityGivenNodeComData(filename string, nodeComList []nodeCom) {
	fmt.Println("Writing file...")
	f, err := os.Create(filename)
	Handle(err)
	defer f.Close()

	w := bufio.NewWriter(f)
	for i, nc := range nodeComList {
		_, err := w.WriteString(nc.node.Name + "\t")
		Handle(err)
		for _, com := range nc.community {
			_, err := w.WriteString(strconv.Itoa(com) + " ")
			Handle(err)
		}
		if i != len(nodeComList)-1 {
			_, err = w.WriteString("\n")
			Handle(err)
		}
	}
	w.Flush()
}

func (g Graph) createNodeCommunityList() []nodeCom {
	fmt.Println("Compiling data...")
	var nodeComPair []nodeCom
	for _, node := range g.Nodes {
		pair := nodeCom{node, nil}
		nodeComPair = append(nodeComPair, pair)
	}
	for i, c := range g.Communities {
		nodelist := c.ListNodes()
		for _, node := range nodelist {
			nodeComPair[node.Index].community = append(nodeComPair[node.Index].community, i+1)
		}
	}
	counter := len(g.Communities)
	for i, c := range nodeComPair {
		if c.community == nil {
			nodeComPair[i].community = append(nodeComPair[i].community, counter)
		}
		counter = counter + 1
	}
	// for _, pair := range nodeComPair {
	// 	fmt.Println(pair.node.Print(), pair.community)
	// }
	return nodeComPair
}
