package gosolve

import (
	"fmt"
)

const (
	NO_SOLUTION = iota
	SOLVED      = iota
)

type Solution struct {
	status   int8
	solution []int8
}

type SearchParams struct {
	max_depth int8
	max_sols  int8
	move_set  *[]Move
	pt        *PruningTable
}

func ida_star_worker(params *SearchParams, p *Puzzle, current_sol *[]int8,
	depth int8, n_sols *int8, sols *[][]int8) {
	if depth < params.max_depth {
		new_ps := p.ApplyMoves(params.move_set)
		length := int8(len(*params.move_set))
		for i := int8(0); i < length; i++ {
			h := (*params.pt).Lookup(&new_ps[i])
			if h+depth > params.max_depth {
				continue
			}
			(*current_sol)[depth] = i
			ida_star_worker(params, &new_ps[i], current_sol, depth+1, n_sols, sols)
		}
	} else {
		if p.IsSolved() {
			(*sols)[*n_sols] = make([]int8, len(*current_sol))
			copy((*sols)[*n_sols], *current_sol)
			*n_sols++
			fmt.Println(*n_sols)
		}
	}
	if *n_sols == params.max_sols {
		return
	}
}

func IDA_Star(pt *PruningTable, p *Puzzle, move_set *[]Move, max_sols int8,
	max_depth int8) [][]int8 {
	sols := make([][]int8, max_sols)
	n_sols := int8(0)
	for i := int8(0); i <= max_depth; i++ {
		current_sol := make([]int8, i)
		params := SearchParams{i, max_sols, move_set, pt}
		ida_star_worker(&params, p, &current_sol, 0, &n_sols, &sols)
		if n_sols == max_sols {
			break
		}
	}
	return sols
}
