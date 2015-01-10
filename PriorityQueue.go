package gosolve

const MAX_PQ_SIZE = 100

type Node struct {
	State         Puzzle
	Moves_applied []int
}

type pqnode struct {
	node     Node
	priority int8
}

type PriorityQueue struct {
	queue [MAX_PQ_SIZE]pqnode
	size  int
}

func NewPriorityQueue() PriorityQueue {
	return PriorityQueue{[MAX_PQ_SIZE]pqnode{}, 0}
}

func (pq *PriorityQueue) shift_right(index int) {
	for i := pq.size - 1; i >= index; i-- {
		pq.queue[i+1] = pq.queue[i]
	}
	pq.size++
}

func (pq *PriorityQueue) Insert(n Node, priority int8) {
	pqn := pqnode{n, priority}
	if pq.size == 0 { // empty
		pq.queue[0] = pqn
		pq.size++
		return
	}
	for i := 0; i < pq.size; i++ {
		if pq.queue[i].priority < priority { // insert in middle of queue
			pq.shift_right(i)
			pq.queue[i] = pqn
			return
		}
	}
	pq.queue[pq.size] = pqn // insert at end
	pq.size++
}

func (pq *PriorityQueue) Pop() Node {
	pq.size--
	return pq.queue[pq.size].node
}

func (pq *PriorityQueue) IsEmpty() bool {
	if pq.size > 0 {
		return false
	} else {
		return true
	}
}

func (pq *PriorityQueue) IsFull() bool {
	if pq.size == MAX_PQ_SIZE {
		return true
	} else {
		return false
	}
}
