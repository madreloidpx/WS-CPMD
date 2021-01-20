package main

import (
	"flag"
	"fmt"
)

// How to run
// go build wcpmd.go node.go edge.go wq.go graph.go
// go run wcpmd -fn tests/10/network.dat -out results/10/community.dat

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

	graph := CreateGraphByFile(filename)
	graph.CreateWeakCliques()
	graph.FormCommunity()
	//graph.Print()
	graph.GenerateCommunityFile(outfile)
	//generate nmi
	//show graph
	fmt.Println("Done.")
}

//Handle panics when there's an error
func Handle(err error) {
	if err != nil {
		panic(err)
	}
}
