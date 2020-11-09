package graphs

import (
	"testing"
)

// TestGraphCreate is responsible for testing the initialization of the graph
func TestGraphCreate(t *testing.T) {
	// TODO: use testing package
	graph := Graph{}
	graph.Init()
	graph.Insert("A")
	graph.Insert("B")
	graph.Insert("C")
	graph.Connect("A", "B", 1)
	graph.Connect("A", "C", 100)
	graph.Connect("B", "C", -1)
	graph.Connect("C", "B", 1)
	graph.Represent()
}
