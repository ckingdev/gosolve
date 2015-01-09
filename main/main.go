package main

import (
	"fmt"
	"github.com/cpalone/gosolve"
)

func main() {
	htm := gosolve.GetHTMMoves()
	p := gosolve.GetSolvedPuzzle()
	pt := gosolve.GetPruningTable(p, 6, htm, 3000000)
	p = p.Apply(&htm[1])
	p = p.Apply(&htm[4])
	p = p.Apply(&htm[1])
	sols := gosolve.IDA_Star(&pt, &p, &htm, 2, 4)
	fmt.Println(sols)
}
