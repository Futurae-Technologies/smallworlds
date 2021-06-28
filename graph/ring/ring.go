package ring

import (
	"math"
	"math/rand"
	"time"

	"futurae.com/smallworlds/graph"
)

// Graph holds the generated nodes and edges according to
// k and beta parameters. K controls how many short edges
// within the ring are created, while beta controls how the
// short edges are rewired.
type Graph struct {
	n     int
	k     int
	beta  float64
	nodes []int
	edges map[int]map[int]struct{}
	rand  *rand.Rand
}

// NewGraph creates an empty graph with k set to kOver2 * 2.
// This ensures that k is an even int.
func NewGraph(kOver2 int, beta float64) *Graph {
	return &Graph{
		n:     0,
		k:     kOver2 * 2,
		beta:  beta,
		nodes: make([]int, 0),
		edges: make(map[int]map[int]struct{}),
		rand:  rand.New(rand.NewSource(time.Now().UTC().UnixNano())),
	}
}

// WithSeed is a builder that sets the random number generator.
func (w *Graph) WithSeed(seed int64) *Graph {
	w.rand = rand.New(rand.NewSource(seed))

	return w
}

// WithNodes is a builder that adds _n_ new nodes to the graph.
func (w *Graph) WithNodes(n int) *Graph {
	w.n = w.n + n
	for i := 0; i < n; i++ {
		w.nodes = append(w.nodes, i)
	}

	return w
}

// WithShortEdges is a builder that sets up k short edges per node.
// The edges set up a ring structure. Nodes are initially arranged
// in a list, then each node is linked to its immediate neighbours,
// with the tail of list linking to the head of the list.
func (w *Graph) WithShortEdges() *Graph {
	for _, u := range w.nodes {
		for next := u - (w.k / 2); next < u+(w.k/2+1); next++ {
			if u != next {
				w.addEdge(u, (w.n+next)%w.n)
			}
		}
	}

	return w
}

// WithDistantEdges is a builder that removes short edges given the beta parameter
// and replaces it with a random edge.
func (w *Graph) WithDistantEdges() *Graph {
	for _, u := range w.nodes {
		for next := u - (w.k / 2); next < u+(w.k/2+1); next++ {

			if (u != next) && w.hasEdge(u, (w.n+next)%w.n) {
				if w.rand.Float64() < w.beta {

					w.removeEdge(u, (w.n+next)%w.n)

					added := false
					for !added {
						p := w.rand.Intn(w.n)
						if !w.hasEdge(u, p) && (p != u) {
							w.addEdge(u, p)
							added = true
						}
					}
				}
			}
		}
	}

	return w
}

// Nodes exports the ring elements as graph nodes.
func (w *Graph) Nodes() []graph.Node {
	ns := make([]graph.Node, 0)
	for k := range w.nodes {
		ns = append(ns, graph.IntNode(k))
	}
	return ns
}

// Edges exports the edges as the slice of graph edges.
func (w *Graph) Edges() []graph.Edge {
	es := make([]graph.Edge, 0)

	for from := range w.edges {
		for to := range w.edges[from] {
			es = append(es, graph.IntEdge{graph.IntNode(from), graph.IntNode(to)})
		}
	}
	return es
}

// Valid returns true if n >> k >> ln(n) >> 1.
// A graph needs to be valid before adding long edges in order to guarantee
// the small world property.
func (w *Graph) Valid() bool {
	return (w.n >= 5*w.k) && (float64(w.k) >= 5*math.Log(float64(w.n)))
}

func (w *Graph) hasEdge(from int, to int) bool {
	_, ok := w.edges[from]
	if !ok {
		return false
	}

	_, ok = w.edges[from][to]
	return ok
}

func (w *Graph) addEdge(p, q int) {
	w.addDiEdge(p, q)
	w.addDiEdge(q, p)
}

func (w *Graph) removeEdge(p, q int) {
	w.removeDiEdge(p, q)
	w.removeDiEdge(q, p)
}

func (w *Graph) addDiEdge(from int, to int) {
	_, ok := w.edges[from]
	if !ok {
		w.edges[from] = make(map[int]struct{})
	}

	w.edges[from][to] = struct{}{}
}

func (w *Graph) removeDiEdge(from int, to int) {
	_, ok := w.edges[from]
	if !ok {
		return
	}

	_, ok = w.edges[from][to]
	if !ok {
		return
	}
	delete(w.edges[from], to)
}
