package gosolve

import (
	"fmt"
)

type PruningTable map[Puzzle]int8

func dls(pt *PruningTable, p *Puzzle, depth int, max_depth int, move_set []Move) {
	if depth > 0 {
		// _, ok := pt[p]
		// if ok {
		// 	return
		// }
		new_ps := p.ApplyMoves(move_set)
		for _, new_p := range new_ps {
			dls(pt, &new_p, depth-1, max_depth, move_set)
		}
	} else {
		_, ok := (*pt)[*p]
		if !ok {
			(*pt)[*p] = int8(max_depth)
		}
	}
}

func GetPruningTable(p Puzzle, max_depth int, move_set []Move) PruningTable {
	pt := make(PruningTable)
	for d := 0; d <= max_depth; d++ {
		dls(&pt, &p, d, d, move_set)
		fmt.Printf("Depth: %v: %v\n", d, len(pt))
	}
	return pt
}
