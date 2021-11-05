package world

import (
	"math"
	"math/rand"
	"time"

	"futurae.com/smallworlds/graph"
)

// World type defines a contextual graph (aka a world)
// for agents to traverse.
type World struct {
	toInt    map[string]int
	toNode   map[int]graph.Node
	array    [][]int
	n        int
	rand     *rand.Rand
	contexts map[string]Context
}

// NewWorld creates a world from a given graph. The created world
// has an empty context at each node. Internally the graph is represented
// as an integer matrix to faciliatate k shortest path computations as agents
// traverse the world.
func NewWorld(w graph.Graph) *World {
	toInt := make(map[string]int)
	toNode := make(map[int]graph.Node)
	contexts := make(map[string]Context)

	for i, node := range w.Nodes() {
		toInt[node.String()] = i
		toNode[i] = node
		contexts[node.String()] = NewContext()
	}
	n := len(toInt)

	array := make([][]int, n, n)
	for i := 0; i < n; i++ {
		array[i] = make([]int, n, n)
	}

	for _, edge := range w.Edges() {
		array[toInt[edge.From().String()]][toInt[edge.To().String()]] = 1
	}

	return &World{
		n:        n,
		array:    array,
		toInt:    toInt,
		toNode:   toNode,
		contexts: contexts,
		rand:     rand.New(rand.NewSource(time.Now().UTC().UnixNano())),
	}
}

// WithSeed is a builder that sets the random number generator.
func (w *World) WithSeed(seed int64) *World {
	w.rand = rand.New(rand.NewSource(seed))

	return w
}

// AddEdges inserts the given edges into the world.
func (m *World) AddEdges(es []graph.Edge) {
	for _, edge := range es {
		m.array[m.toInt[edge.From().String()]][m.toInt[edge.To().String()]] = 1
	}
}

// RemoveEdges deletes the given edges from the world.
func (m *World) RemoveEdges(es []graph.Edge) {
	for _, edge := range es {
		m.array[m.toInt[edge.From().String()]][m.toInt[edge.To().String()]] = 0
	}
}

// HasEdge returns true if there is an edge between the given nodes in the graph.
func (m *World) HasEdge(from graph.Node, to graph.Node) bool {
	return m.hasEdge(m.toInt[from.String()], m.toInt[to.String()])
}

func (m *World) hasEdge(from, to int) bool {
	return m.array[from][to] == 1
}

// AddContext sets the context c for the given node.
func (w *World) AddContext(n graph.Node, c Context) *World {
	w.contexts[n.String()] = c
	return w
}

// AddContextWithSpread sets the context c for the given node and its nearest neighbours.
func (w *World) AddContextWithSpread(origin graph.Node, c Context, spread int) {
	w.AddContext(origin, c)
	if spread>0 {
		for _, n := range w.Neighbourhood(origin) {
			w.AddContextWithSpread(n, c, spread-1)
		}
	}
}

// AddIfNotExistsContextFeature adds the feaure (key, value) to the node's context
// if that key is not present.
func (w *World) AddIfKeyNotExists(n graph.Node, key string, value float64) *World {
	w.contexts[n.String()].LeftJoin(Context{key: value})

	return w
}

// Context returns the node's current context.
func (w *World) Context(node graph.Node) Context {
	return w.contexts[node.String()]
}

// Contexts maps the given walk onto a list of respective contexts.
func (w *World) Contexts(walk Walk) []Context {
	ctxs := make([]Context, len(walk), len(walk))
	for i := 0; i < len(walk); i++ {
		ctxs[i] = w.contexts[walk[i].String()]
	}

	return ctxs
}

// Nodes returns all nodes in the graph.
func (w *World) Nodes() []graph.Node {
	ns := make([]graph.Node, 0, 0)
	for _, n := range w.toNode {
		ns = append(ns, n)
	}
	return ns
}

// Edges returns all the edges in the graph.
func (w *World) Edges() []graph.Edge {
	es := make([]graph.Edge, 0, len(w.array))
	for i := 0; i < len(w.array); i++ {
		for j := 0; j < len(w.array); j++ {
			if w.hasEdge(i, j) {
				es = append(es, graph.TupleEdge{w.toNode[i], w.toNode[j]})
			}
		}
	}
	return es
}

// Neighbourhood returns all immediate neighbours of the given node.
func (m *World) Neighbourhood(n graph.Node) []graph.Node {
	return m.toNodes(m.neighbourhood(m.toInt[n.String()]))
}

func (m *World) neighbourhood(n int) []int {
	ns := make([]int, 0, 0)
	for i, p := range m.array[n] {
		if p == 1 {
			ns = append(ns, i)
		}
	}

	return ns
}

func (m *World) toNodes(p path) []graph.Node {
	ns := make([]graph.Node, len(p), len(p))

	for i := 0; i < len(p); i++ {
		ns[i] = m.toNode[p[i]]
	}

	return ns
}

// AvgClusteringCoeff returns the average clustering coefficient for the world.
func (m *World) AvgClusteringCoeff() float64 {
	sum := 0.0
	l := 0

	for _, v := range m.LocalClusteringCoeffs() {
		sum += v
		l++
	}
	return sum / float64(l)
}

// LocalClusteringCoeffs returns the list of local clustering coefficients for each node.
func (m *World) LocalClusteringCoeffs() []float64 {
	localCoefficients := make([]float64, 0, 0)

	for i := 0; i < m.n; i++ {
		localCoefficients = append(localCoefficients, m.localClusteringCoeffFor(i))
	}
	return localCoefficients
}

// LocalClusteringCoeffFor returns the local clustering coefficient for the given node.
func (m *World) LocalClusteringCoeffFor(n graph.Node) float64 {
	return m.localClusteringCoeffFor(m.toInt[n.String()])
}

func (m *World) localClusteringCoeffFor(n int) float64 {
	ns := m.neighbourhood(n)
	if len(ns) == 0 {
		return 0.0
	}

	ns = append(ns, n)
	numLinks := 0

	for _, n1 := range ns {
		for _, n2 := range ns {
			if m.hasEdge(n1, n2) {
				numLinks++
			}
		}
	}

	return float64(numLinks) / float64((len(ns) * (len(ns) - 1)))
}

// RandomWalk returns a random walk of the given length starting at the given node.
func (m *World) RandomWalk(length int, from graph.Node) Walk {
	return m.toNodes(m.randomPath(length, m.toInt[from.String()]))
}

// KShortestPaths computes at most _k_ shortest paths between the given nodes.
// It implements Dijkstra's algorithm.
func (m *World) KShortestPaths(k int, from graph.Node, to graph.Node) []Walk {
	ps := m.kShortestPaths(k, m.toInt[from.String()], m.toInt[to.String()])
	ns := make([]Walk, len(ps), len(ps))

	for i := 0; i < len(ps); i++ {
		ns[i] = m.toNodes(ps[i])
	}

	if len(ps) == 0 {
		return []Walk{{from}}
	}
	return ns
}

// ShortestPathsLens returns the lenght of shortest paths between all nodes.
// It implements Floyd-Warshall algorithm.
func (m *World) ShortestPathsLens() [][]int {
	distArray := m.initializeShortestPaths()

	for k := 0; k < len(distArray); k++ {
		for i := 0; i < len(distArray); i++ {
			for j := 0; j < len(distArray); j++ {

				if distArray[i][j] > distArray[i][k]+distArray[k][j] {
					distArray[i][j] = distArray[i][k] + distArray[k][j]
				}
			}
		}
	}

	return distArray
}

func (m *World) initializeShortestPaths() [][]int {
	distArray := make([][]int, 0, 0)

	for i := 0; i < len(m.array); i++ {
		row := make([]int, 0, 0)
		for j := 0; j < len(m.array); j++ {
			if i == j {
				row = append(row, 0)
			} else if m.hasEdge(i, j) {
				row = append(row, 1)
			} else {
				row = append(row, math.MaxUint32)
			}
		}
		distArray = append(distArray, row)
	}
	return distArray
}
