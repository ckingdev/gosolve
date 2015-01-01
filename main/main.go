package main

import (
	"fmt"
	"github.com/cpalone/gosolve"
)

func main() {
	htm := gosolve.GetHTMMoves()
	p := gosolve.GetSolvedPuzzle()
	fmt.Println(p.IsSolved())
	p = p.Apply(htm[1])
	fmt.Println(p.IsSolved())
	p = p.Apply(htm[1])
	fmt.Println(p.IsSolved())
	p = p.Apply(htm[1])
	fmt.Println(p.IsSolved())
	p = p.Apply(htm[1])
	fmt.Println(p.IsSolved())

}
