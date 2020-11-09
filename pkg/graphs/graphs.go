package graphs

import (
	"errors"
	"fmt"
)

var (
	// ErrNodeDoesNotExist is used when an operation is performed on a node which is not present in the graph
	ErrNodeDoesNotExist = errors.New("graphs: operation performed on a Node which does not exist in Graph")
)

// Node is an interface that represents any item that can be put into the graph
type Node interface {
}

// Graph implements a directed weighted graph using an adjacency matrix using nested hashmaps
type Graph struct {
	adjacencyMatrix map[Node]map[Node]int
}

// Init initializes the graph
func (g *Graph) Init() {
	g.adjacencyMatrix = make(map[Node]map[Node]int)
}

// Insert inserts a node into the graph
func (g *Graph) Insert(x Node) {
	g.adjacencyMatrix[x] = make(map[Node]int)
}

// Connect connects Node a to Node b with a weight
func (g *Graph) Connect(a Node, b Node, weight int) error {
	// check if a is present in the adjacency matrix. if not return error
	_, ok := g.adjacencyMatrix[a]
	if !ok {
		return ErrNodeDoesNotExist
	}
	// connect the nodes providing a weight value
	g.adjacencyMatrix[a][b] = weight
	return nil
}

// Represent outputs a human readable form of graph to stdout
// TODO: complete function
func (g *Graph) Represent() {
	fmt.Print(g.adjacencyMatrix)
}

// DFS performs a depth first search of the graph from node A to B. returns slice of path as well as cost. nil, -1 if no path found
func (g *Graph) DFS(a Node, b Node) ([]Node, int) {
	return nil, 0
}
