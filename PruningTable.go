package gosolve

import (
	"fmt"
)

type PruningTable struct {
	max_depth int8
	table     map[Puzzle]int8
}

func PTLookup(pt *PruningTable, p *Puzzle) int8 {
	val, ok := pt.table[*p]
	if ok {
		return val
	} else {
		return pt.max_depth + 1
	}
}

func dls(pt *PruningTable, p *Puzzle, depth int8, max_depth int8, move_set []Move) {
	if depth > 0 {
		new_ps := p.ApplyMoves(move_set)
		for _, new_p := range new_ps {
			dls(pt, &new_p, depth-1, max_depth, move_set)
		}
	} else {
		_, ok := pt.table[*p]
		if !ok {
			pt.table[*p] = int8(max_depth)
		}
	}
}

func GetPruningTable(p Puzzle, max_depth int8, move_set []Move) PruningTable {
	pt := PruningTable{max_depth, make(map[Puzzle]int8)}
	for d := int8(0); d <= max_depth; d++ {
		dls(&pt, &p, d, d, move_set)
		fmt.Printf("Depth: %v: %v\n", d, len(pt.table))
	}
	return pt
}
