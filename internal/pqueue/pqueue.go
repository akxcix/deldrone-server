package pqueue

// Item is struct which holds values of any type in our priority queue
type Item struct {
	Value    interface{}
	Priority int
}

// PQueue is the priority queue of Items
type PQueue []Item

// Len returns length of the PQueue
func (pq PQueue) Len() int {
	return len(pq)
}

// Less defines priority. Lower numerical value of Item.priority means higher priority
func (pq PQueue) Less(i, j int) bool {
	return pq[i].Priority < pq[j].Priority
}

// Swap items in the PQueue
func (pq PQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

// Push Item to PQueue
func (pq *PQueue) Push(x interface{}) {
	*pq = append(*pq, x.(Item))
}

// Pop an Item from PQueue
func (pq *PQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

// IsEmpty returns true only when pq.Len() == 0
func (pq PQueue) IsEmpty() bool {
	if pq.Len() == 0 {
		return true
	}
	return false
}
