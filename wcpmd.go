package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
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
	printCommunityData(outfile, communityConverted, unassigned)
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

func printCommunityData(filename string, community [][]string, nonmember []string) {
	f, err := os.Create(filename)
	Handle(err)
	defer f.Close()

	mappedCommunity := make(map[string][]string)
	for i, c := range community {
		for _, node := range c {
			v, exist := mappedCommunity[node]
			if exist == false {
				var comIn []string
				comIn = append(comIn, strconv.Itoa((i + 1)))
				mappedCommunity[node] = comIn
			} else {
				v = append(v, strconv.Itoa((i + 1)))
				mappedCommunity[node] = v
			}

		}
	}

	fmt.Println(mappedCommunity)

	w := bufio.NewWriter(f)
	for k, v := range mappedCommunity {
		_, err := w.WriteString(k + "\t")
		Handle(err)
		for i, comm := range v {
			_, err := w.WriteString(comm)
			Handle(err)
			if i < len(comm) {
				_, err := w.WriteString(" ")
				Handle(err)
			}
		}
		_, err = w.WriteString("\n")
		Handle(err)
	}
	for i, nm := range nonmember {
		_, err := w.WriteString(nm)
		Handle(err)
		if i < len(nonmember)-1 {
			_, err := w.WriteString("\n")
			Handle(err)
		}
	}
	w.Flush()
}

func Handle(err error) {
	if err != nil {
		panic(err)
	}
}
