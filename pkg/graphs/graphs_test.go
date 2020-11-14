package graphs

import (
	"fmt"
	"sort"
	"testing"
)

// helper function to check whether two slices are equal
func checkEqualNodeSlice(actual []Node, want []Node, t *testing.T) {
	if actual == nil && want != nil {
		t.Errorf("want %v got %v", want, actual)
	} else if actual != nil && want == nil {
		t.Errorf("want %v got %v", want, actual)
	} else if actual != nil && want != nil {
		if len(actual) != len(want) {
			t.Errorf("want %v got %v", want, actual)
		}
		for i, v := range want {
			if v != actual[i] {
				t.Errorf("want %v got %v", want, actual)
			}
		}
	}
}

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

	for i, tt := range tests {
		testName := fmt.Sprintf("%d", i)
		t.Run(testName, func(t *testing.T) {
			actual := g.Connect(tt.a, tt.b, tt.w)
			if actual != tt.want {
				t.Errorf("want %v got %v", tt.want, actual)
			}

		})
	}
}

func TestGraphConnectBoth(t *testing.T) {
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
		{1, 3, -1, ErrNegativeEdgeWeight},
		{1, 3, 1, nil},
	}

	for i, tt := range tests {
		testName := fmt.Sprintf("%d", i)
		t.Run(testName, func(t *testing.T) {
			actual := g.ConnectBoth(tt.a, tt.b, tt.w)
			if actual != tt.want {
				t.Errorf("want %v got %v", tt.want, actual)
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
			checkEqualNodeSlice(actual, tt.want, t)
		})
	}
}

func TestGraphAll(t *testing.T) {
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
		want []Node
	}{
		{[]Node{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("all")
		t.Run(testName, func(t *testing.T) {
			actual := g.All()

			sort.Slice(actual, func(i, j int) bool {
				return actual[i].(int) < actual[j].(int)
			})

			sort.Slice(tt.want, func(i, j int) bool {
				return tt.want[i].(int) < tt.want[j].(int)
			})

			checkEqualNodeSlice(actual, tt.want, t)

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
		a       Node
		want    []Node
		wantErr error
	}{
		{"A", []Node{"A", "E", "C", "F", "B", "D"}, nil},
		{"G", nil, ErrNodeDoesNotExist},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("%v", tt.a)
		t.Run(testName, func(t *testing.T) {
			actual, actualErr := g.DFS(tt.a)
			if actualErr != tt.wantErr {
				t.Errorf("want %v got %v", tt.wantErr, actualErr)
			}
			checkEqualNodeSlice(actual, tt.want, t)
		})
	}
}

func TestGraphDijkstra(t *testing.T) {
	g := Graph{}
	g.Init()

	g.Insert("A")
	g.Insert("B")
	g.Insert("C")
	g.Insert("D")
	g.Insert("E")
	g.Insert("F")
	g.Insert("G")

	g.Connect("A", "B", 1)
	g.Connect("A", "E", 2)
	g.Connect("B", "C", 3)
	g.Connect("B", "D", 4)
	g.Connect("C", "A", 5)
	g.Connect("C", "F", 6)
	g.Connect("D", "F", 7)
	g.Connect("E", "C", 8)
	g.Connect("F", "E", 9)

	tests := []struct {
		a         Node
		b         Node
		wantSlice []Node
		wantCost  int
		wantErr   error
	}{
		{"A", "F", []Node{"F", "C", "B", "A"}, 10, nil},
		{"C", "D", []Node{"D", "B", "A", "C"}, 10, nil},
		{"A", "G", nil, 0, ErrUnreachableNode},
		{"A", "I", nil, 0, ErrNodeDoesNotExist},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("%s to %s", tt.a, tt.b)
		t.Run(testName, func(t *testing.T) {
			actualSlice, actualCost, actualErr := g.Dijkstra(tt.a, tt.b)
			if actualErr != tt.wantErr {
				t.Errorf("want %v got %v", tt.wantErr, actualErr)
			}
			if actualCost != tt.wantCost {
				t.Errorf("want %v got %v", tt.wantCost, actualCost)
			}
			checkEqualNodeSlice(actualSlice, tt.wantSlice, t)
		})
	}
}
