package main

import (
	"fmt"
	"log"

	"github.com/timohahaa/ETU_term3_algorithms_and_data_structures/trees"
)

func main() {
	treeString := "(8 (9 (5)) (1))"
	tree, err := trees.BinaryTreeFromBrackets(treeString)
	if err != nil {
		log.Fatal(err)
	}
	levels := tree.LevelOrderTraversal()
	for _, level := range levels {
		fmt.Println(level)
	}
}
