package random

import (
	"math/rand"
	"time"

	"futurae.com/smallworlds/graph"
)

// Graph holds randomly generated nodes
// and edges
type Graph struct {
	nodes []struct{}
	edges map[int]map[int]struct{}
	rand  *rand.Rand
}

// NewGraph creates an empty graph
// with a rand seed set to time.Now()
func NewGraph() *Graph {
	return &Graph{
		nodes: make([]struct{}, 0, 0),
		edges: make(map[int]map[int]struct{}),
		rand:  rand.New(rand.NewSource(time.Now().UTC().UnixNano())),
	}
}

// WithSeed is a builder function that sets
// the random number generator.
func (w *Graph) WithSeed(seed int64) *Graph {
	w.rand = rand.New(rand.NewSource(seed))

	return w
}

// WithNodes is a builder function that adds
// _n_ nodes to the graph.
func (w *Graph) WithNodes(n int) *Graph {
	for i := 0; i < n; i++ {
		w.nodes = append(w.nodes, struct{}{})
	}

	return w
}

// WithEdges is a builder function that adds
// _n_ new edges between randomly chones nodes.
func (w *Graph) WithEdges(n int) *Graph {
	bound := len(w.nodes)
	added := 0

	for {
		from := w.rand.Intn(bound)
		to := w.rand.Intn(bound)

		if (from != to) && !(w.hasEdge(from, to)) {
			w.addEdge(from, to)
			added++
		}

		if added >= n {
			break
		}
	}

	return w
}

// Nodes exports the internal slice representing nodes
// to a slice of IntNodes
func (w *Graph) Nodes() []graph.Node {
	ns := make([]graph.Node, 0, 0)
	for k := range w.nodes {
		ns = append(ns, graph.IntNode(k))
	}
	return ns
}

// Edges exports all internal edges as a slice of IntEdges.
func (w *Graph) Edges() []graph.Edge {
	es := make([]graph.Edge, 0, 0)

	for from := range w.edges {
		for to := range w.edges[from] {
			es = append(es, graph.IntEdge{graph.IntNode(from), graph.IntNode(to)})
		}
	}
	return es
}

func (w *Graph) hasEdge(from int, to int) bool {
	_, ok := w.edges[from]
	if !ok {
		return false
	}

	_, ok = w.edges[from][to]
	return ok
}

func (w *Graph) addEdge(from int, to int) {
	_, ok := w.edges[from]
	if !ok {
		w.edges[from] = make(map[int]struct{})
	}

	w.edges[from][to] = struct{}{}
}
