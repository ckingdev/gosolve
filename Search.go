package gosolve

const (
	NO_SOLUTION    = iota
	ALREADY_SOLVED = iota
	SOLVED         = iota
)

type Solution struct {
	status   int8
	solution []int8
}

type SearchParams struct {
	max_depth int8
	max_sols  int8
	n_sols    *int8
	move_set  *[]Move
	pt        *PruningTable
	sols      *[][]int8
}

func ida_star_worker(params *SearchParams, p *Puzzle, current_sol *[]int8,
	depth int8) {
	if depth < params.max_depth {
		new_ps := p.ApplyMoves(params.move_set)
		length := int8(len(*params.move_set))
		for i := int8(0); i < length; i++ {
			h := (*params.pt).Lookup(&new_ps[i])
			if h+depth > params.max_depth {
				continue
			}
			(*current_sol)[depth] = i
			ida_star_worker(params, p, current_sol, depth+1)
			if *params.n_sols == params.max_sols {
				return
			}
		}
	} else {
		if p.IsSolved() {
			(*params.sols)[*params.n_sols] = *current_sol
			*params.n_sols++
		}
	}
}
