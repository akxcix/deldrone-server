package graphs

import (
	"fmt"
	"testing"
)

// TestGraphInsert is responsible for testing Insertion of elements in the graph
func TestGraphInsert(t *testing.T) {
	// TODO: use testing package
	var tests = []int{
		1,
		10,
		100,
		1_000,
		10_000,
		100_000,
		1_000_000,
		10_000_000,
	}
	for _, tt := range tests {
		testName := fmt.Sprintf("%d", tt)
		t.Run(testName, func(t *testing.T) {
			g := new(Graph)
			g.Init()
			for i := 0; i < tt; i++ {
				g.Insert(i)
			}
		})
	}
}

func TestGraphConnect(t *testing.T) {
	g := Graph{}
	g.Init()
	for i := 0; i < 10; i++ {
		g.Insert(i)
	}
	var tests = []struct {
		a    Node
		b    Node
		w    int
		want error
	}{
		{1, 2, 1, nil},
		{-1, 2, 1, ErrNodeDoesNotExist},
		{1, -1, 1, ErrNodeDoesNotExist},
		{1, 9, 0, nil},
		{2, 1, 100, nil},
		{1, 1, 1, nil},
		{1, 3, 1, nil},
		{1, 3, -1, ErrNegativeEdgeWeight},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("%d, %d, %d, %e", tt.a, tt.b, tt.w, tt.want)
		t.Run(testName, func(t *testing.T) {
			actual := g.Connect(tt.a, tt.b, tt.w)
			if actual != tt.want {
				t.Errorf("want %e got %e", tt.want, actual)
			}

		})
	}
}

func TestGraphNeighbours(t *testing.T) {
	g := Graph{}
	g.Init()
	for i := 0; i < 10; i++ {
		g.Insert(i)
	}
	g.Connect(1, 2, 1)
	g.Connect(1, 3, 7)
	g.Connect(5, 4, 1)
	g.Connect(4, 2, 8)
	g.Connect(15, 1, 4)

	tests := []struct {
		a    Node
		want []Node
	}{
		{1, []Node{2, 3}},
		{2, []Node{}},
		{3, []Node{}},
		{4, []Node{2}},
		{5, []Node{4}},
		{6, []Node{}},
		{7, []Node{}},
		{8, []Node{}},
		{9, []Node{}},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("%d", tt.a)
		t.Run(testName, func(t *testing.T) {
			actual := g.Neighbours(tt.a)
			if actual == nil && tt.want != nil {
				t.Errorf("want %t got %t", tt.want, actual)
			} else if actual != nil && tt.want == nil {
				t.Errorf("want %t got %t", tt.want, actual)
			} else if actual != nil && tt.want != nil {
				if len(actual) != len(tt.want) {
					t.Errorf("want %t got %t", tt.want, actual)
				}
				for i, v := range tt.want {
					if v != actual[i] {
						t.Errorf("want %t got %t", tt.want, actual)
					}
				}
			}
		})
	}
}

func TestGraphDFS(t *testing.T) {
	g := Graph{}
	g.Init()

	g.Insert("A")
	g.Insert("B")
	g.Insert("C")
	g.Insert("D")
	g.Insert("E")
	g.Insert("F")

	g.Connect("A", "B", 1)
	g.Connect("A", "E", 1)
	g.Connect("B", "C", 1)
	g.Connect("B", "D", 1)
	g.Connect("C", "A", 1)
	g.Connect("C", "F", 1)
	g.Connect("D", "F", 1)
	g.Connect("E", "C", 1)
	g.Connect("F", "E", 1)

	tests := []struct {
		a    Node
		want []Node
	}{
		{"A", []Node{"A", "E", "C", "F", "B", "D"}},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("%t", tt.a)
		t.Run(testName, func(t *testing.T) {
			actual := g.DFS("A")
			if actual == nil && tt.want != nil {
				t.Errorf("want %t got %t", tt.want, actual)
			} else if actual != nil && tt.want == nil {
				t.Errorf("want %t got %t", tt.want, actual)
			} else if actual != nil && tt.want != nil {
				if len(actual) != len(tt.want) {
					t.Errorf("want %t got %t", tt.want, actual)
				}
				for i, v := range tt.want {
					if v != actual[i] {
						t.Errorf("want %t got %t", tt.want, actual)
					}
				}
			}
		})
	}
}
