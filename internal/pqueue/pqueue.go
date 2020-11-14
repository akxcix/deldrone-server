package pqueue

// Item is struct which holds values of any type in our priority queue
type Item struct {
	value    interface{}
	priority int
}

// PQueue is the priority queue of Items
type PQueue []Item

// Len returns length of the PQueue
func (pq PQueue) Len() int {
	return len(pq)
}

// Less defines priority. Lower numerical value of Item.priority means higher priority
func (pq PQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
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

// // Item is what a priority queue would be storing
// type Item struct {
// 	value    interface{} // value of the item
// 	priority int         // priority
// 	index    int         // index in the heap. managed by heap.Interface
// }

// // PriorityQueue is a slice holding the refrences to Items
// type PriorityQueue []*Item

// // Len returns length of priority queue
// func (pq PriorityQueue) Len() int {
// 	return len(pq)
// }

// // Less is responsible for Pop to provide no. with high priority (lower Item.priority value)
// func (pq PriorityQueue) Less(i, j int) bool {
// 	return pq[i].priority < pq[j].priority
// }

// func (pq PriorityQueue) Swap(i, j int) {
// 	pq[i], pq[j] = pq[j], pq[i]
// 	pq[i].index = i
// 	pq[j].index = j
// }

// // Push pushes an item to the pqueue at the last position
// func (pq *PriorityQueue) Push(x interface{}) {
// 	n := len(*pq)
// 	item := x.(*Item)
// 	item.index = n
// 	*pq = append(*pq, item)
// }

// // Pop removes the first item from the pqueue
// func (pq *PriorityQueue) Pop() interface{} {
// 	old := *pq
// 	n := len(old)
// 	item := old[n-1]
// 	old = old[0 : n-1]
// 	return item
// }

// func (pq *PriorityQueue) update(item *Item, value interface{}, priority int) {
// 	item.value = value
// 	item.priority = priority
// 	heap.Fix(pq, item.index)
// }
