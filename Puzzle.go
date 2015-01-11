// Package gosolve provides functionality to work with 2x2 puzzles.
// Puzzle, move, and pruning table generation and manipulation are currently
// implemented, as is searching for solved states with IDA*.
package gosolve

// Location encodes a position of a piece, in terms of its permutation p and
// orientation o.
type Location struct {
	p int8
	o int8
}

// Piece encodes a piece, in terms of its id (essentially, where the piece
// should be while solved) and its location (where it is currently).
type Piece struct {
	id  int8
	loc Location
}

// Puzzle is a collection of 8 pieces.
type Puzzle [8]Piece

// Move is a map of id -> permutation & twist.
type Move [8]Location

// NewMove returns the identity Move.
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

func getBaseMoves() [3]Move {
	var baseMoves [3]Move
	for i := 0; i < 3; i++ {
		baseMoves[i] = NewMove()
	}
	baseMoves[0][0] = Location{3, 1}
	baseMoves[0][3] = Location{7, 2}
	baseMoves[0][7] = Location{4, 1}
	baseMoves[0][4] = Location{0, 2}

	baseMoves[1][2] = Location{3, 0}
	baseMoves[1][3] = Location{0, 0}
	baseMoves[1][0] = Location{1, 0}
	baseMoves[1][1] = Location{2, 0}

	baseMoves[2][1] = Location{0, 1}
	baseMoves[2][0] = Location{4, 2}
	baseMoves[2][4] = Location{5, 1}
	baseMoves[2][5] = Location{1, 2}
	return baseMoves
}

// GetHTMMoves returns a slice of the moves in <R, U, F> in half-turn metric.
func GetHTMMoves() []Move {
	baseMoves := getBaseMoves()
	htmMoves := make([]Move, 9)
	for i := 0; i < 3; i++ {
		htmMoves[i*3] = baseMoves[i]
		htmMoves[i*3+1] = compose(baseMoves[i], baseMoves[i])
		htmMoves[i*3+2] = compose(baseMoves[i], htmMoves[i*3+1])

	}
	return htmMoves
}

// GetSolvedPuzzle returns a solved Puzzle.
func GetSolvedPuzzle() Puzzle {
	var p Puzzle
	for i := 0; i < 8; i++ {
		p[i] = Piece{int8(i), Location{int8(i), 0}}
	}
	return p
}

// Apply applies a Move to a Puzzle and returns the result.
func (p *Puzzle) Apply(m *Move) Puzzle {
	var newP Puzzle
	for i := 0; i < 8; i++ {
		tmp := (*m)[p[i].loc.p]
		newP[i].loc = Location{tmp.p, (tmp.o + p[i].loc.o) % 3}
		newP[i].id = p[i].id
	}
	return newP
}

// ApplyMoves returns a slice of the Puzzles resulting from applying each
// move to the Puzzle independently.
func (p *Puzzle) ApplyMoves(moveSet *[]Move) []Puzzle {
	appliedPs := make([]Puzzle, len(*moveSet))
	for i, move := range *moveSet {
		appliedPs[i] = p.Apply(&move)
	}
	return appliedPs
}

// ApplySequence applies a given sequence of moves to a Puzzle in series.
func (p *Puzzle) ApplySequence(moves []int, moveSet *[]Move) Puzzle {
	newP := *p
	for i := range moves {
		newP = newP.Apply(&(*moveSet)[i])
	}
	return newP
}

// IsSolved returns true if the Puzzle is solved, otherwise false.
func (p *Puzzle) IsSolved() bool {
	for i := 0; i < 8; i++ {
		if p[i].loc.p != p[i].id || p[i].loc.o != 0 {
			return false
		}
	}
	return true
}
