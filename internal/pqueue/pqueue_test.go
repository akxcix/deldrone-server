package pqueue

import (
	"container/heap"
	"testing"
)

func TestPQueue(t *testing.T) {
	// TODO: write robust tests
	pq := make(PQueue, 0)
	items := []Item{
		Item{
			Value:    "A",
			Priority: 1,
		},
		Item{
			Value:    1,
			Priority: 2,
		},
		Item{
			Value:    "B",
			Priority: 3,
		},
		Item{
			Value:    "C",
			Priority: 3,
		},
		Item{
			Value:    "D",
			Priority: 4,
		},
		Item{
			Value:    "E",
			Priority: 5,
		},
		Item{
			Value:    "F",
			Priority: 0,
		},
	}

	heap.Init(&pq)
	for _, item := range items {
		heap.Push(&pq, item)
		t.Log(pq)
	}

	for !pq.IsEmpty() {
		t.Log(heap.Pop(&pq))
	}

}
