package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	var filename string
	flag.StringVar(&filename, "fn", "", "ex test.dat")

	var outfile string
	flag.StringVar(&outfile, "out", "community.dat", "ex community.dat")

	flag.Parse()

	if filename == "" {
		fmt.Println("Please provide a filename.")
		return
	}

	data := ReadNodeFile(filename)
	nodes := GenerateNodeData(data)
	fmt.Println("Node Data:")
	fmt.Println(nodes)
	weakcliques := GenerateWeakCliqueData(nodes)
	fmt.Println("Weak Clique Data:")
	fmt.Println(weakcliques)
	communities := weakcliques.GenerateCommunityData()
	fmt.Println("Communities:")
	var communityConverted [][]string
	for _, v := range communities {
		var convertedMembers []string
		for member := range v.Members {
			val, _ := nodes.IndexMap[member]
			convertedMembers = append(convertedMembers, val.Name)
		}
		communityConverted = append(communityConverted, convertedMembers)
	}
	fmt.Println(communityConverted)
	var unassigned []string
	for _, v := range nodes.ShowNodes() {
		marked := false
		for _, cm := range communityConverted {
			for _, c := range cm {
				if v == c {
					marked = true
				}
			}
		}
		if marked == false {
			unassigned = append(unassigned, v)
		}
	}
	fmt.Println("Nodes with no communities:")
	fmt.Println(unassigned)
}

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

func Handle(err error) {
	if err != nil {
		panic(err)
	}
}

// func printCommunityGivenNodeComData(filename string, nodeComList []nodeCom) {
// 	fmt.Println("Writing file...")
// 	f, err := os.Create(filename)
// 	Handle(err)
// 	defer f.Close()

// 	w := bufio.NewWriter(f)
// 	for i, nc := range nodeComList {
// 		_, err := w.WriteString(nc.node.Name + "\t")
// 		Handle(err)
// 		for _, com := range nc.community {
// 			_, err := w.WriteString(strconv.Itoa(com) + " ")
// 			Handle(err)
// 		}
// 		if i != len(nodeComList)-1 {
// 			_, err = w.WriteString("\n")
// 			Handle(err)
// 		}
// 	}
// 	w.Flush()
// }
