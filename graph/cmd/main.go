package main

import (
	"fmt"
	"log"

	"github.com/timohahaa/ETU_term3_algorithms_and_data_structures/graph"
)

func main() {
	g, err := graph.NewGraph("text.txt")
	if err != nil {
		fmt.Println("Error while initing a graph:")
		log.Fatal(err)
	}

	fmt.Println("Initial graph:")
	for i := 0; i < g.Edges.Len(); i++ {
		fmt.Println(g.Edges.GetNoError(i))
	}
	fmt.Println()

	mst := g.ComputeMST()
	sum := 0
	fmt.Println("MST:")
	for i := 0; i < mst.Len(); i++ {
		e := mst.GetNoError(i)
		sum += e.Weight
		fmt.Println(e)
	}

	fmt.Println()
	fmt.Printf("Total MST weight: %v\n", sum)
}
