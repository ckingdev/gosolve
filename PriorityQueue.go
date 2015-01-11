package gosolve

const maxPQSize = 100

// Node encodes a Puzzle and the moves taken from the search root to arrive
// at the puzzle.
type Node struct {
	State        Puzzle
	MovesApplied []int
}

type pqnode struct {
	node     Node
	priority int8
}

// PriorityQueue is a min-priority queue that holds Nodes.
type PriorityQueue struct {
	queue [maxPQSize]pqnode
	size  int
}

// NewPriorityQueue returns a new, empty PriorityQueue.
func NewPriorityQueue() PriorityQueue {
	return PriorityQueue{[maxPQSize]pqnode{}, 0}
}

func (pq *PriorityQueue) shiftRight(index int) {
	for i := pq.size - 1; i >= index; i-- {
		pq.queue[i+1] = pq.queue[i]
	}
	pq.size++
}

// Insert inserts a Node into the PriorityQueue in sorted order.
func (pq *PriorityQueue) Insert(n Node, priority int8) {
	pqn := pqnode{n, priority}
	if pq.size == 0 { // empty
		pq.queue[0] = pqn
		pq.size++
		return
	}
	for i := 0; i < pq.size; i++ {
		if pq.queue[i].priority < priority { // insert in middle of queue
			pq.shiftRight(i)
			pq.queue[i] = pqn
			return
		}
	}
	pq.queue[pq.size] = pqn // insert at end
	pq.size++
}

// Pop removes the Node with lowest priority from the queue and returns it.
func (pq *PriorityQueue) Pop() Node {
	pq.size--
	return pq.queue[pq.size].node
}

// IsEmpty checks whether the queue is empty or not.
func (pq *PriorityQueue) IsEmpty() bool {
	if pq.size > 0 {
		return false
	}
	return true
}

// IsFull checks whether the queue is full or not.
func (pq *PriorityQueue) IsFull() bool {
	if pq.size == maxPQSize {
		return true
	}
	return false
}
