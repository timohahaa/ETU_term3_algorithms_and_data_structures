package main

import (
	"lds"
)

func main() {
	l := lds.NewDLList[int]()
	l.AddHead(8)
	l.AddHead(7)
	l.AddHead(1337)
	l.Print()
}
