package main

import (
	"fmt"
	"lds"
	"log"
)

func main() {
	//	l := lds.NewDLList[int]()
	//	l.AddHead(8)
	//	l.AddHead(7)
	//	l.AddHead(1337)
	//	l.Print()

	str, err := lds.SortingStation("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(str)
}
