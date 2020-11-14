package graphs

// TODO: Rewrite. Works rn but can be optimized a lot

import (
	"container/heap"
	"errors"
	"math"

	"github.com/iamadarshk/deldrone-server/internal/pqueue"
)

var (
	// ErrNodeDoesNotExist is used when an operation is performed on a node which is not present
	// in the graph
	ErrNodeDoesNotExist = errors.New("graphs: operation performed on a Node which does not exist in Graph")

	// ErrNegativeEdgeWeight is used when the graph is provided a negative edge weight
	ErrNegativeEdgeWeight = errors.New("graphs: negative edge weight in graph")

	// ErrUnreachableNode is used when a pathfinding algorithm isn't able to reach destination node
	ErrUnreachableNode = errors.New("graphs: the destination node is unreachable")
)

var (
	// +Inf
	inf = math.MaxInt64
)

// Node is an empty interface that represents any item that can be put into the graph
type Node interface{}

// PathNode is a container used to hold the predecessor of the current node through which the
// cost of traversal is minimized
type PathNode struct {
	contained Node
	via       Node
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

// ConnectBoth connects a to b and b to a with weight
func (g *Graph) ConnectBoth(a, b Node, weight int) error {
	err := g.Connect(a, b, weight)
	if err != nil {
		return err
	}
	err = g.Connect(b, a, weight)
	if err != nil {
		return err
	}
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

// All returns list of all nodes that are present in the graph
func (g Graph) All() []Node {
	all := make([]Node, len(g.adjacencyMatrix))
	i := 0
	for key := range g.adjacencyMatrix {
		all[i] = key
		i++
	}
	return all
}

// Exists returns true if the node exists in the graph
func (g *Graph) Exists(a Node) bool {
	return false
}

// DFS performs a DFS on the graph and returns slice of Nodes in the order they were found during
// the search
func (g *Graph) DFS(a Node) ([]Node, error) {
	_, ok := g.adjacencyMatrix[a]
	if !ok {
		return nil, ErrNodeDoesNotExist
	}

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
	return found, nil
}

// Dijkstra provides the shortest path and its cost in a graph from a to b. returns
// error if it's not possible to reach
func (g *Graph) Dijkstra(source, destination Node) ([]Node, int, error) {
	// check if both nodes exist in graph
	_, ok1 := g.adjacencyMatrix[source]
	_, ok2 := g.adjacencyMatrix[destination]
	if !ok1 || !ok2 {
		return nil, 0, ErrNodeDoesNotExist
	}

	pq := pqueue.PQueue{}      // pq maintains order based on cost
	heap.Init(&pq)             // init pq
	visited := map[Node]bool{} // keeps track of node we have already visited
	all := g.All()             // get slice of all nodes
	processed := []pqueue.Item{}

	// add all nodes to pq
	for _, node := range all { // range over all the nodes in the graph
		if node == source {
			// if the node is source node, set priority 0 and predecessor nil
			heap.Push(&pq, pqueue.Item{
				Value: PathNode{
					contained: source,
					via:       nil,
				},
				Priority: 0,
			})
		} else {
			// otherwise set priority = infinity, predecessor = source
			heap.Push(&pq, pqueue.Item{
				Value: PathNode{
					contained: node,
					via:       source,
				},
				Priority: inf,
			})
		}
	}

	curr := pq[0] // var curr holds the node which is being processed
	for {
		currConatained := curr.Value.(PathNode).contained
		// if we have already visited the node, skip
		if visited[currConatained] == true {
			curr = heap.Pop(&pq).(pqueue.Item)
			continue
		}

		// else compute distance for each node. if it''s less, update
		neighbours := g.Neighbours(currConatained)
		for _, neighbour := range neighbours {
			for i := 0; i < pq.Len(); i++ {
				other := pq[i].Value.(PathNode).contained
				if neighbour != other {
					continue
				}
				if curr.Priority+g.adjacencyMatrix[currConatained][other] < pq[i].Priority {
					// update pq if a shorter path is found
					pq[i] = pqueue.Item{
						Value: PathNode{
							contained: other,
							via:       currConatained,
						},
						Priority: curr.Priority + g.adjacencyMatrix[currConatained][other],
					}
					heap.Init(&pq) // fix order of pq
				}
			}
		}

		visited[currConatained] = true      // mark node as visited
		processed = append(processed, curr) // copy processed node to a different location

		// break out if destination is reached
		if currConatained == destination {
			break
		}

		curr = heap.Pop(&pq).(pqueue.Item)
	}

	pathMap := map[Node]Node{} // maps a node to its predecessor in the shortest path
	costMap := map[Node]int{}  // maps a node to the cost of the path required to reach it

	// populate pathMap and costMap
	for _, pathNode := range processed {
		node := pathNode.Value.(PathNode).contained
		pathMap[node] = pathNode.Value.(PathNode).via
		costMap[node] = pathNode.Priority
	}

	// return error if cost == inf
	if costMap[destination] == inf {
		return nil, 0, ErrUnreachableNode
	}

	path := []Node{}                 // holds the shortest path
	path = append(path, destination) // append destination to list
	prev := pathMap[destination]     // store it's predecessor
	for prev != nil {                // while predecessor to current node exists
		path = append(path, prev) // add that predecessor
		prev = pathMap[prev]      // update predecessor
	}

	cost := costMap[destination] // find cost to reach destination

	return path, cost, nil
}
