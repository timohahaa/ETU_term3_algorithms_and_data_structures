package main

import (
	//	"fmt"

	"fmt"

	"github.com/timohahaa/ETU_term3_algorithms_and_data_structures/timsort"
)

// if testing using user input, GC kicks in and the time will be slower + not accurate at all
func main() {
	//	var size int
	//	fmt.Print("size of an array to test on: ")
	//	fmt.Scanln(&size)
	for i := 1; i <= 10; i++ {
		timsort.Test(i * 100)
		fmt.Println()
	}
}
