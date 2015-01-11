package gosolve

import (
	"fmt"
)

type searchParams struct {
	maxDepth int8
	maxSols  int8
	moveSet  *[]Move
	pt       *PruningTable
}

func idaStarWorker(params *searchParams, p *Puzzle, currentSol *[]int8,
	depth int8, nSols *int8, sols *[][]int8) {
	if depth < params.maxDepth {
		newPs := p.ApplyMoves(params.moveSet)
		length := int8(len(*params.moveSet))
		for i := int8(0); i < length; i++ {
			h := (*params.pt).Lookup(&newPs[i])
			if h+depth > params.maxDepth {
				continue
			}
			(*currentSol)[depth] = i
			idaStarWorker(params, &newPs[i], currentSol, depth+1, nSols, sols)
		}
	} else {
		if p.IsSolved() {
			(*sols)[*nSols] = make([]int8, len(*currentSol))
			copy((*sols)[*nSols], *currentSol)
			*nSols++
			fmt.Println(*nSols)
		}
	}
	if *nSols == params.maxSols {
		return
	}
}

// IDAStar uses the given PruningTable to solve a puzzle with the IDA* search
// algorithm. Returns at most maxSols solutions.
func IDAStar(pt *PruningTable, p *Puzzle, moveSet *[]Move, maxSols int8,
	maxDepth int8) [][]int8 {
	sols := make([][]int8, maxSols)
	nSols := int8(0)
	for i := int8(0); i <= maxDepth; i++ {
		currentSol := make([]int8, i)
		params := searchParams{i, maxSols, moveSet, pt}
		idaStarWorker(&params, p, &currentSol, 0, &nSols, &sols)
		if nSols == maxSols {
			break
		}
	}
	return sols
}
