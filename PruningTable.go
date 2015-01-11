package gosolve

import (
	"fmt"
)

// PruningTable stores and retrieves heuristic values for informed search.
type PruningTable struct {
	maxDepth int8
	table    map[Puzzle]int8
}

// Lookup retrieves the heuristic value for a given puzzle.
func (pt *PruningTable) Lookup(p *Puzzle) int8 {
	val, ok := pt.table[*p]
	if ok {
		return val
	}
	return pt.maxDepth + 1
}

func dls(pt *PruningTable, p *Puzzle, depth int8, maxDepth int8, moveSet *[]Move) {
	if depth > 0 {
		newPs := p.ApplyMoves(moveSet)
		for _, newP := range newPs {
			dls(pt, &newP, depth-1, maxDepth, moveSet)
		}
	} else {
		_, ok := pt.table[*p]
		if !ok {
			pt.table[*p] = int8(maxDepth)
		}
	}
}

// GetPruningTable generates a pruning table for the given puzzle to a
// specified depth and returns it.
func GetPruningTable(p Puzzle, maxDepth int8, moveSet []Move, sizeHint int) PruningTable {
	pt := PruningTable{maxDepth, make(map[Puzzle]int8, sizeHint)}
	for d := int8(0); d <= maxDepth; d++ {
		dls(&pt, &p, d, d, &moveSet)
		fmt.Printf("Depth: %v: %v\n", d, len(pt.table))
	}
	return pt
}

// GetNumEntries returns the number of entries in the PruningTable.
func (pt *PruningTable) GetNumEntries() int {
	return len(pt.table)
}
