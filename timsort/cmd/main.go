package main

import (
	"fmt"

	"github.com/timohahaa/ETU_term3_algorithms_and_data_structures/timsort"
)

func main() {
	arr := []int{4, 3, 2, 10, 12, 1, 5, 6}
	timsort.Test(arr)
	fmt.Println(arr)
}
