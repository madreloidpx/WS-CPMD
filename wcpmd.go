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
