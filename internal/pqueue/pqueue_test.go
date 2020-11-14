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
			value:    "A",
			priority: 1,
		},
		Item{
			value:    1,
			priority: 2,
		},
		Item{
			value:    "B",
			priority: 3,
		},
		Item{
			value:    "C",
			priority: 3,
		},
		Item{
			value:    "D",
			priority: 4,
		},
		Item{
			value:    "E",
			priority: 5,
		},
		Item{
			value:    "F",
			priority: 0,
		},
	}

	heap.Init(&pq)
	for _, item := range items {
		heap.Push(&pq, item)
		t.Log(pq)
	}

	for len(pq) > 0 {
		t.Log(heap.Pop(&pq))
	}

}
