package main

import (
	"fmt"

	"github.com/timohahaa/ETU_term3_algorithms_and_data_structures/trees"
)

func main() {
	t := trees.NewAVLTree[int](func(a, b int) int {
		if a > b {
			return 1
		}
		if a < b {
			return -1
		}
		return 0
	})
	t.Insert(8)
	t.Insert(4)
	t.Insert(2)
	t.Insert(9)
	t.Insert(41)
	t.Insert(14)
	t.Insert(28)
	levels := t.LevelOrderTraversal()
	for _, level := range levels {
		fmt.Println(level)
	}
	fmt.Println()
	fmt.Println(t.InorderTraversal())
	fmt.Println()
	fmt.Println(t.PreorderTraversal())
	fmt.Println()
	fmt.Println(t.PostorderTraversal())
	fmt.Println()
	fmt.Println(t.PostorderTraversalRecursionless())
	fmt.Println()
	t.PrintTree()

}
