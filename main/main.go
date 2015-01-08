package main

import (
	"github.com/cpalone/gosolve"
)

func main() {
	htm := gosolve.GetHTMMoves()
	p := gosolve.GetSolvedPuzzle()
	_ = gosolve.GetPruningTable(p, 6, htm)
}
