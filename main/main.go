package main

import (
	// "encoding/binary"
	// "fmt"
	// "reflect"

	"github.com/cpalone/gosolve"
)

func main() {
	p := gosolve.GetSolvedPuzzle()
	n := gosolve.Node{p, []int{0, 0, 0, 0}}
	pq := gosolve.NewPriorityQueue()
	pq.Insert(n, 1)
	pq.Insert(n, 3)
	pq.Insert(n, 2)
	pq.PrintElement(0)
	pq.PrintElement(1)
	pq.PrintElement(2)
}
