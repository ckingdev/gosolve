package gosolve

import (
	"testing"
)

func TestSolvedPuzzle(t *testing.T) {
	p := GetSolvedPuzzle()
	if !p.IsSolved() {
		t.Fatal("Solved puzzle does not return true for IsSolved.")
	}
}

func TestMoves(t *testing.T) {
	htm := GetHTMMoves()
	for i := 0; i < 9; i += 3 {
		move := htm[i]
		p := GetSolvedPuzzle()
		p1 := p.Apply(&move)
		if p1.IsSolved() {
			t.Fatalf("Move %v is solved when applied once to a solved puzzle.", i)
		}
		p2 := p1.Apply(&move)
		if p2.IsSolved() {
			t.Fatalf("Move %v is solved when applied twice to a solved puzzle.", i)
		}
		pp := p2.Apply(&move)
		if pp.IsSolved() {
			t.Fatalf("Move %v is solved when applied 3 times to a solved puzzle.", i)
		}
		solved := pp.Apply(&move)
		if !solved.IsSolved() {
			t.Fatalf("Move %v is not solved when applied 4 times to a solved puzzle.", i)
		}
	}
}

func TestGetPruningTable(t *testing.T) {
	htm := GetHTMMoves()
	p := GetSolvedPuzzle()
	pt := GetPruningTable(p, 6, htm, 62360)
	if len(pt.table) != 62360 {
		t.Fatalf("Pruning table has %v entries, %v required.", len(pt.table), 62360)
	}
}

func TestPTLookup(t *testing.T) {
	solved := GetSolvedPuzzle()
	htm := GetHTMMoves()
	pt := GetPruningTable(solved, 2, htm, 100)
	solved_depth := PTLookup(&pt, &solved)
	if solved_depth != 0 {
		t.Fatalf("Solved puzzle has depth %v, expected 0.", solved_depth)
	}
	depth4 := solved.Apply(&htm[0])
	depth4 = depth4.Apply(&htm[3])
	depth4 = depth4.Apply(&htm[6])
	depth4 = depth4.Apply(&htm[0])
	depth4_depth := PTLookup(&pt, &depth4)
	if depth4_depth != 3 {
		t.Fatalf("Incorrect depth for puzzle not in table, got %v, expected %v.", depth4_depth, 3)
	}
}

func BenchmarkGetPruningTable(b *testing.B) {
	htm := GetHTMMoves()
	p := GetSolvedPuzzle()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = GetPruningTable(p, 6, htm, 62360)
	}
}
