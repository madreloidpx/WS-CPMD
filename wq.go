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
}

type WeakCliqueFactory struct {
	WeakCliques map[string]WeakClique
	Length      int
}

func WeakCliques() WeakCliqueFactory {
	wqf := WeakCliqueFactory{
		WeakCliques: make(map[string]WeakClique),
		Length:      0,
	}
	return wqf
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

func (wqf *WeakCliqueFactory) CreateWeakClique(nodes []Node) {
	nodeList := make([]int, 0, len(nodes))
	for _, n := range nodes {
		nodeList = append(nodeList, n.Index)
	}
	hexname := HexName(nodeList)
	_, exist := wqf.WeakCliques[hexname]
	if exist == false {
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
		wqf.WeakCliques[hexname] = wq
		wqf.Length = wqf.Length + 1
	}
}

func (wqf *WeakCliqueFactory) MergeCliques(hexname string, wq1, wq2 WeakClique) {
	members := mergeMap(wq1.Members, wq2.Members)
	inlink := mergeMap(wq1.InLink, wq2.InLink)
	outlink := mergeMap(wq1.OutLink, wq2.OutLink)
	wq := WeakClique{
		Members: members,
		InLink:  inlink,
		OutLink: outlink,
		Merged:  false,
	}
	wqf.WeakCliques[hexname] = wq
	wqf.Length = wqf.Length + 1
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

func GenerateWeakCliqueData(nodes NodeFactory) WeakCliqueFactory {
	weakcliques := WeakCliques()
	for nodeIndex, node := range nodes.IndexMap {
		for outIndex := range node.OutLink {
			outNode := nodes.IndexMap[outIndex]
			for inIndex := range node.InLink {
				if nodeIndex != outIndex && nodeIndex != inIndex && outIndex != inIndex {
					linked, _ := outNode.OutLink[inIndex]
					if linked == true {
						inNode := nodes.IndexMap[inIndex]
						nodeList := []Node{node, outNode, inNode}
						weakcliques.CreateWeakClique(nodeList)
					}
				}
			}
			for otherOutIndex := range node.OutLink {
				if nodeIndex != outIndex && nodeIndex != otherOutIndex && outIndex != otherOutIndex {
					linked, _ := outNode.OutLink[otherOutIndex]
					if linked == true {
						otherOutNode := nodes.IndexMap[otherOutIndex]
						nodeList := []Node{node, outNode, otherOutNode}
						weakcliques.CreateWeakClique(nodeList)
					}
				}
			}
		}
	}
	return weakcliques
}

func (wq *WeakClique) MarkMerged() {
	wq.Merged = true
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

func (wqf WeakCliqueFactory) GenerateCommunityData() map[string]WeakClique {
	community := make(map[string]WeakClique)
	wqData := wqf.WeakCliques
	for {
		tmpWQData := MergeCliquesIteration(wqData)
		for wqHex, wq := range wqData {
			if wq.Merged == false {
				fmt.Println("Adding to community:", wq)
				community[wqHex] = wq
			}
		}
		wqData = tmpWQData
		if len(wqData) == 0 {
			break
		}
	}
	community2 := MergeCliquesIteration(community)
	for wqHex, wq := range community {
		if wq.Merged == false {
			fmt.Println("Adding to community:", wq)
			community2[wqHex] = wq
		}
	}
	return community2
}

func MergeCliquesIteration(wqData map[string]WeakClique) map[string]WeakClique {
	tmpWQData := WeakCliques()
	for i, wq1 := range wqData {
		for j, wq2 := range wqData {
			if i != j {
				members := mergeMap(wq1.Members, wq2.Members)
				memlist := ShowMembers(members)
				hexName := HexName(memlist)
				_, exist := tmpWQData.WeakCliques[hexName]
				if exist == true {
					continue
				}
				sameIn := CompareLinks(wq1.InLink, wq2.InLink)
				sameIn = CompareLinks(sameIn, members)
				sameOut := CompareLinks(wq1.OutLink, wq2.OutLink)
				sameOut = CompareLinks(sameOut, members)
				if len(members) > len(sameIn)+1 || len(members) > len(sameOut)+1 {
					continue
				}
				fmt.Println("Merging", wq1, wq2)
				tmpWQData.MergeCliques(hexName, wq1, wq2)
				wq1.MarkMerged()
				wq2.MarkMerged()
				wqData[i] = wq1
				wqData[j] = wq2
			}
		}
	}
	return tmpWQData.WeakCliques
}
