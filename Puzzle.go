package gosolve

type Location struct {
	p int8
	o int8
}

type Piece struct {
	id  int8
	loc Location
}

type Puzzle [8]Piece

type Move [8]Location

func NewMove() Move {
	var m Move
	for i := 0; i < 8; i++ {
		m[i] = Location{int8(i), 0}
	}
	return m
}

func compose(a Move, b Move) Move {
	ab := NewMove()
	for i := 0; i < 8; i++ {
		tmp := a[i]
		tmp = b[tmp.p]
		tmp.o = (tmp.o + a[i].o) % 3
		ab[i] = tmp
	}
	return ab
}

func get_base_moves() [3]Move {
	var base_moves [3]Move
	for i := 0; i < 3; i++ {
		base_moves[i] = NewMove()
	}
	base_moves[0][0] = Location{3, 1}
	base_moves[0][3] = Location{7, 2}
	base_moves[0][7] = Location{4, 1}
	base_moves[0][4] = Location{0, 2}

	base_moves[1][2] = Location{3, 0}
	base_moves[1][3] = Location{0, 0}
	base_moves[1][0] = Location{1, 0}
	base_moves[1][1] = Location{2, 0}

	base_moves[2][1] = Location{0, 1}
	base_moves[2][0] = Location{4, 2}
	base_moves[2][4] = Location{5, 1}
	base_moves[2][5] = Location{1, 2}
	return base_moves
}

func GetHTMMoves() []Move {
	base_moves := get_base_moves()
	htm_moves := make([]Move, 9)
	for i := 0; i < 3; i++ {
		htm_moves[i*3] = base_moves[i]
		htm_moves[i*3+1] = compose(base_moves[i], base_moves[i])
		htm_moves[i*3+2] = compose(base_moves[i], htm_moves[i*3+1])

	}
	return htm_moves
}

func GetSolvedPuzzle() Puzzle {
	var p Puzzle
	for i := 0; i < 8; i++ {
		p[i] = Piece{int8(i), Location{int8(i), 0}}
	}
	return p
}

func (p *Puzzle) Apply(m *Move) Puzzle {
	var new_p Puzzle
	for i := 0; i < 8; i++ {
		tmp := (*m)[p[i].loc.p]
		new_p[i].loc = Location{tmp.p, (tmp.o + p[i].loc.o) % 3}
		new_p[i].id = p[i].id
	}
	return new_p
}

func (p *Puzzle) ApplyMoves(move_set *[]Move) []Puzzle {
	applied_ps := make([]Puzzle, len(*move_set))
	for i, move := range *move_set {
		applied_ps[i] = p.Apply(&move)
	}
	return applied_ps
}

func (p *Puzzle) IsSolved() bool {
	for i := 0; i < 8; i++ {
		if p[i].loc.p != p[i].id || p[i].loc.o != 0 {
			return false
		}
	}
	return true
}
