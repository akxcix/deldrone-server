package graphs

import (
	"errors"
	"fmt"
)

var (
	// ErrNodeDoesNotExist is used when an operation is performed on a node which is not present
	// in the graph
	ErrNodeDoesNotExist = errors.New("graphs: operation performed on a Node which does not exist in Graph")

	// ErrNegativeEdgeWeight is used when the graph is provided a negative edge weight
	ErrNegativeEdgeWeight = errors.New("graphs: negative edge weight in graph")
)

// Node is an empty interface that represents any item that can be put into the graph
type Node interface{}

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
	// check if negative egde weight is provided to graph. If so, return error
	if weight < 0 {
		return ErrNegativeEdgeWeight
	}

	// check if a is present in the adjacency matrix. if not return error
	_, ok1 := g.adjacencyMatrix[a]
	_, ok2 := g.adjacencyMatrix[b]
	if !ok1 || !ok2 {
		return ErrNodeDoesNotExist
	}
	g.adjacencyMatrix[a][b] = weight // connect the nodes providing a weight value
	return nil
}

// Neighbours gets neighbours of a node. specifically the nodes you can go to from node a.
func (g *Graph) Neighbours(a Node) []Node {
	neighbours := make([]Node, len(g.adjacencyMatrix[a]))
	i := 0
	for key := range g.adjacencyMatrix[a] {
		neighbours[i] = key
		i++
	}
	return neighbours
}

// Exists returns true if the node exists in the graph
func (g *Graph) Exists(a Node) bool {
	return false
}

// Represent outputs a human readable form of graph to stdout
// TODO: complete function
func (g *Graph) Represent() {
	fmt.Print(g.adjacencyMatrix)
}

// DFS performs a DFS on the graph and returns slice of Nodes in the order they were found during
// the search
func (g *Graph) DFS(a Node) []Node {
	visited := map[Node]bool{}
	found := []Node{}
	stack := []Node{}
	curr := a
	visited[curr] = true
	found = append(found, curr)
	for _, neighbour := range g.Neighbours(curr) {
		stack = append(stack, neighbour)
	}
	for len(stack) != 0 {
		curr, stack = stack[len(stack)-1], stack[:len(stack)-1] // pop
		if visited[curr] == true {
			continue
		}
		visited[curr] = true
		found = append(found, curr)
		for _, neighbour := range g.Neighbours(curr) {
			stack = append(stack, neighbour)
		}
	}
	return found
}
