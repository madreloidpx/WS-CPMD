package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	nodes := Nodes()
	nodes.ShowNodes()
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
	fmt.Printf(err.Error())
}
