package main

import (
	"fmt"
	"math"
	"sort"
)

type WeakClique struct {
	Members map[int]bool
	InLink  map[int]bool
	OutLink map[int]bool
	Merged  bool
	Parent  *Node
}

func mergeMap(m1, m2 map[int]bool) map[int]bool {
	merged := make(map[int]bool)
	for k, v := range m1 {
		merged[k] = v
	}
	for k, v := range m2 {
		merged[k] = v
	}
	return merged
}

func CreateWeakClique(nodes []*Node) WeakClique {
	members := make(map[int]bool)
	inlink := make(map[int]bool)
	outlink := make(map[int]bool)
	for i, n := range nodes {
		members[n.Index] = true
		if i == 0 {
			inlink = n.InLink
			outlink = n.OutLink
		} else {
			inlink = mergeMap(inlink, n.InLink)
			outlink = mergeMap(outlink, n.OutLink)
		}
	}
	wq := WeakClique{
		Members: members,
		InLink:  inlink,
		OutLink: outlink,
		Merged:  false,
	}
	return wq
}

func MergeCliques(wq1, wq2 *WeakClique) WeakClique {
	members := mergeMap(wq1.Members, wq2.Members)
	inlink := mergeMap(wq1.InLink, wq2.InLink)
	outlink := mergeMap(wq1.OutLink, wq2.OutLink)
	wq := WeakClique{
		Members: members,
		InLink:  inlink,
		OutLink: outlink,
		Merged:  false,
		Parent:  nil,
	}
	return wq
}

func (wq WeakClique) ShowMembers() []int {
	cliques := make([]int, 0, len(wq.Members))
	for k := range wq.Members {
		cliques = append(cliques, k)
	}
	return cliques
}

func ShowMembers(members map[int]bool) []int {
	cliques := make([]int, 0, len(members))
	for k := range members {
		cliques = append(cliques, k)
	}
	return cliques
}

func HexName(members []int) string {
	hexMap := make(map[int]int)
	for _, member := range members {
		hexIndex := member / 4
		hexPos := member % 4
		hexValue := math.Pow(2, float64(hexPos))
		currVal := hexMap[hexIndex]
		hexMap[hexIndex] = currVal + int(hexValue)
	}
	hexList := make([]int, 0, len(hexMap))
	for k := range hexMap {
		hexList = append(hexList, k)
	}
	sort.Ints(hexList)
	name := ""
	for i := 0; i <= hexList[len(hexList)-1]; i++ {
		name = fmt.Sprintf("%X", hexMap[i]) + name
	}
	return name
}

func (nodes *NodeFactory) GenerateWeakCliqueData() {
	for nodeIndex, node := range nodes.IndexMap {
		fmt.Println("Generating WQ data for node:", nodeIndex)
		for outIndex := range node.OutLink {
			nodeMembers := make([]*Node, 0)
			outNode, _ := nodes.IndexMap[outIndex]
			nodeMembers = append(nodeMembers, node)
			nodeMembers = append(nodeMembers, outNode) //append first two nodes of weakclique

			for inIndex := range node.InLink { //check all incoming nodes for circular clique
				if nodeIndex != outIndex && nodeIndex != inIndex && outIndex != inIndex {
					_, linked := outNode.OutLink[inIndex]
					if linked == true {
						inNode := nodes.IndexMap[inIndex]
						nodeMembers = append(nodeMembers, inNode)
					}
				}
			}
			for otherOutIndex := range node.OutLink { //check all outgoing again for hierchical clique
				if nodeIndex != outIndex && nodeIndex != otherOutIndex && outIndex != otherOutIndex {
					_, linkedIn := outNode.InLink[otherOutIndex]
					_, linkedOut := outNode.OutLink[otherOutIndex]
					if linkedIn == true || linkedOut == true {
						otherOutNode := nodes.IndexMap[otherOutIndex]
						nodeMembers = append(nodeMembers, otherOutNode)
					}
				}
			}

			if len(nodeMembers) > 2 { //members are not just the compared edge
				nodeList := make([]int, 0, len(nodeMembers))
				for _, n := range nodeMembers {
					nodeList = append(nodeList, n.Index)
				}
				hexname := HexName(nodeList)
				weakClique := CreateWeakClique(nodeMembers)
				node.AddWeakClique(hexname, &weakClique)
			}
		}
		fmt.Println("Weak Cliques:", len(node.WeakCliques))
	}
}

func (wq *WeakClique) MarkMerged() {
	wq.Merged = true
	wq.Parent.AddWQMark()
}

func CompareLinks(list1, list2 map[int]bool) map[int]bool {
	similar := make(map[int]bool)
	for k := range list1 {
		val, _ := list2[k]
		if val == true {
			similar[k] = true
		}
	}
	return similar
}

func (nf NodeFactory) GenerateCommunityData() map[string]*WeakClique {
	for nodeIndex, node := range nf.IndexMap {
		// fmt.Println("In node", node.Name)
		if node.Complete == true {
			// fmt.Println("Already marked complete")
			continue
		}
		for outIndex := range node.OutLink {
			if nodeIndex == outIndex {
				continue
			}
			neighbor := nf.IndexMap[outIndex]
			// fmt.Println("Checking neighbor", neighbor.Name)
			if neighbor.Complete == true {
				// fmt.Println("Neighbor already marked complete")
				continue
			}
			// fmt.Println("Checking for duplicates...")
			mergedWQList := MergeWeakCliqueList(node.WeakCliques, neighbor.WeakCliques)
			// fmt.Println("Merging cliques...")
			MergeCliquesIteration(mergedWQList)
		}
	}
	community := nf.checkCommunityInList()
	lenCom := len(community)
	for {
		MergeCliquesIteration(community)
		community = nf.checkCommunityInList()
		if lenCom == len(community) {
			return community
		}
		lenCom = len(community)
	}
}

func (nf *NodeFactory) checkCommunityInList() map[string]*WeakClique {
	community := make(map[string]*WeakClique)
	for _, node := range nf.IndexMap {
		if node.Complete == false {
			// fmt.Println("Not marked Complete", node)
			for hexname, wq := range node.WeakCliques {
				if wq.Merged == false {
					// fmt.Println("Not marked merged", *wq)
					community[hexname] = wq
				}
			}
		}
	}
	return community
}

func MergeCliquesIteration(wqlist map[string]*WeakClique) bool {
	if len(wqlist) == 0 {
		return true
	}
	var currentWQ *WeakClique
	var currWQHex string
	for i, wq := range wqlist {
		if currentWQ == nil {
			currentWQ = wq
			currWQHex = i
			continue
		}
		members := mergeMap(currentWQ.Members, wq.Members)
		memlist := ShowMembers(members)
		hexName := HexName(memlist)
		_, exist := wqlist[hexName]
		if exist == true {
			if currWQHex == hexName {
				wq.MarkMerged()
				delete(wqlist, i)
			} else if i == hexName {
				currentWQ.MarkMerged()
				delete(wqlist, currWQHex)
				MergeCliquesIteration(wqlist)
				return true
			} else {
				wq.MarkMerged()
				currentWQ.MarkMerged()
				delete(wqlist, i)
				delete(wqlist, currWQHex)
				MergeCliquesIteration(wqlist)
				return true
			}
			continue
		}
		merge := CheckMerge(currentWQ, wq)
		if merge == true {
			wqMerged := MergeCliques(currentWQ, wq)
			currentWQ.Parent.AddWeakClique(hexName, &wqMerged)
			currentWQ.MarkMerged()
			wq.MarkMerged()
			delete(wqlist, i)
			delete(wqlist, currWQHex)
			MergeCliquesIteration(wqlist)
			return true
		}
	}
	delete(wqlist, currWQHex)
	MergeCliquesIteration(wqlist)
	return true
}

func MergeWeakCliqueList(wq1, wq2 map[string]*WeakClique) map[string]*WeakClique {
	mergedList := make(map[string]*WeakClique)
	for i, _wq1 := range wq1 {
		if _wq1.Merged == true {
			continue
		}
		mergedList[i] = _wq1
	}
	for j, _wq2 := range wq2 {
		if _wq2.Merged == true {
			continue
		}
		_, exist := mergedList[j]
		if exist == true {
			_wq2.MarkMerged()
		} else {
			mergedList[j] = _wq2
		}
	}
	return mergedList
}

func CheckMerge(wq1, wq2 *WeakClique) bool {
	members := mergeMap(wq1.Members, wq2.Members)
	sameIn := CompareLinks(wq1.InLink, wq2.InLink)
	sameIn = CompareLinks(sameIn, members)
	sameOut := CompareLinks(wq1.OutLink, wq2.OutLink)
	sameOut = CompareLinks(sameOut, members)
	if len(members) > len(sameIn)+1 || len(members) > len(sameOut)+1 {
		return false
	}
	return true
}
