package main

import "fmt"

//Node object
type Node struct {
	Index       int
	Name        string
	Inlink      []*Edge
	inlinkSize  int32
	Outlink     []*Edge
	outlinkSize int32
}

//CreateNode create node object
func CreateNode(index int, name string) Node {
	n := Node{
		Index:       index,
		Name:        name,
		inlinkSize:  0,
		outlinkSize: 0,
	}
	return n
}

//InsertInlink inserts an edge to incomming link array
func (n *Node) InsertInlink(edge *Edge) {
	n.Inlink = append(n.Inlink, edge)
	n.inlinkSize = n.inlinkSize + 1
}

//InsertOutlink inserts an edge to outgoing link array
func (n *Node) InsertOutlink(edge *Edge) {
	n.Outlink = append(n.Outlink, edge)
	n.outlinkSize = n.outlinkSize + 1
}

//SearchNode searches a node in a node array
func SearchNode(nodes []*Node, n string) int {
	for i, node := range nodes {
		if node.Name == n {
			return i
		}
	}
	return -1
}

//Print print a node
func (n Node) Print() string {
	var in string
	var out string
	for i, edge := range n.Inlink {
		if i == 0 {
			in = edge.Name
		} else {
			in = in + ", " + edge.Name
		}
	}
	for i, edge := range n.Outlink {
		if i == 0 {
			out = edge.Name
		} else {
			out = out + ", " + edge.Name
		}
	}
	return fmt.Sprintf("Node(%s, Out:[%s], In:[%s], Index:%d)", n.Name, out, in, n.Index)
}

//ReduceUniqueNodes reduces a nodelist to unique nodes
func ReduceUniqueNodes(nodelist []*Node, start int) []*Node {
	if len(nodelist) == 0 {
		return nil
	}
	if len(nodelist) == start {
		return nodelist
	}
	unique := nodelist[:start+1]
	for i := start + 1; i < len(nodelist); i++ {
		if nodelist[start] != nodelist[i] {
			unique = append(unique, nodelist[i])
		}
	}
	return ReduceUniqueNodes(unique, start+1)
}

//CompareNodes returns common nodes
func CompareNodes(n1, n2 []*Node) []*Node {
	var commonNodes []*Node
	for _, i := range n1 {
		for _, j := range n2 {
			if i == j {
				commonNodes = append(commonNodes, i)
			}
		}
	}
	return commonNodes
}
