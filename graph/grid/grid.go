package grid

import (
	"math"
	"math/rand"
	"time"

	"futurae.com/smallworlds/graph"
)

// Graph holds the 2d matrix of nodes (hence the Grid model)
// with LenX and LenY dimensions.
//
// Graph nodes are stored in the underlying _positions_ data structure
// that represents a 2d matrix. Graph edges are stored as a set of (from, to)
// tuples. The set data structure is implemented as a map of maps in the _edges_
// structure.
type Graph struct {
	LenX  int
	LenY  int
	nodes positions
	edges edges
	rand  *rand.Rand
}

// NewGraph returns a new empty graph with the given
// maximum dimensions.
func NewGraph(lenX int, lenY int) *Graph {
	return &Graph{
		LenX:  lenX,
		LenY:  lenY,
		nodes: make(positions),
		edges: make(edges),
		rand:  rand.New(rand.NewSource(time.Now().UTC().UnixNano())),
	}
}

// WithSeed is a builder function that sets
// the random number generator.
func (w *Graph) WithSeed(seed int64) *Graph {
	w.rand = rand.New(rand.NewSource(seed))

	return w
}

// WithAllNodes builds a node at every position in the grid.
// I.e. it creates LenX*LenY nodes.
func (w *Graph) WithAllNodes() *Graph {
	for i := 0; i < w.LenX; i++ {
		for j := 0; j < w.LenY; j++ {
			w.addNode(at(i, j))
		}
	}

	return w
}

// WithShortEdges builds an edge between two grid positions if their
// _Manhattan_ distance is less than maxDistance.
func (w *Graph) WithShortEdges(maxDistance int) *Graph {
	for _, from := range w.nodes.slice() {
		adjecantPositions := newPositions().withFromTo(from, maxDistance)

		for _, to := range adjecantPositions.slice() {
			if !from.equal(to) {
				w.addEdgeIfValid(from, to)
			}
		}
	}

	return w
}

// WithDropout randomly drops each edge with the given probability.
func (w *Graph) WithDropout(p float64) *Graph {
	for _, e := range w.edges.slice() {
		if w.rand.Float64() < p {
			w.removeEdge(e)
		}
	}

	return w
}

// WithDistantEdges adds q edges from every node. The ends
// are chosen given the likelihood defined as distance(from, to)^(-1*r).
func (w *Graph) WithDistantEdges(q int, r int) *Graph {
	for _, from := range w.nodes.slice() {
		w.addDistantEdgesFrom(from, q, r)
	}

	return w
}

func (w *Graph) addDistantEdgesFrom(from Position, q int, r int) {
	normConst := w.normalizingConstFor(from, r)
	added := 0

	for added < q {
		for _, to := range w.nodes.slice() {
			if (!from.equal(to)) && (added < q) && (!w.contains(from, to)) {
				p := math.Pow(float64(from.distance(to)), float64(-1*r)) / normConst

				if w.rand.Float64() < p {
					w.addDirectEdgeIfValid(from, to)
					added++
				}
			}
		}
	}
}

// Nodes exports the grid as a slice of Nodes.
func (w *Graph) Nodes() []graph.Node {
	ns := make([]graph.Node, 0, 0)
	for _, n := range w.nodes.slice() {
		ns = append(ns, n)
	}
	return ns
}

// Edges exports the edges as a slice of Edges.
func (w *Graph) Edges() []graph.Edge {
	es := make([]graph.Edge, 0, 0)
	for _, e := range w.edges.slice() {
		es = append(es, edge{e[0], e[1]})
	}
	return es
}

func (w *Graph) normalizingConstFor(node Position, r int) float64 {
	c := 0.0

	for _, d := range w.nodes.distancesFrom(node) {
		c = c + math.Pow(float64(d), float64(-1*r))
	}

	return c
}

func (w *Graph) addNode(p Position) {
	if p.within(w.LenX, w.LenY) {
		w.nodes.add(p)
	}
}

func (w *Graph) addEdgeIfValid(from Position, to Position) {
	if w.exists(from) && w.exists(to) {
		w.edges.add(from, to)
	}
}

func (w *Graph) addDirectEdgeIfValid(from Position, to Position) {
	if w.exists(from) && w.exists(to) {
		w.edges.addDirection(from, to)
	}
}

func (w *Graph) removeEdge(edge []Position) {
	w.edges.remove(edge[0], edge[1])
}

func (w *Graph) exists(p Position) bool {
	return w.nodes.has(p)
}

func (w *Graph) contains(from, to Position) bool {
	return w.edges.contains(from, to)
}
