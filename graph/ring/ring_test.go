package ring

import (
	"testing"

	"futurae.com/smallworlds/graph"
	"github.com/stretchr/testify/assert"
)

func Test_NewGraph(t *testing.T) {
	w := NewGraph(2, 0.0)

	assert.Equal(t, w.n, 0)
	assert.Equal(t, w.k, 4)
	assert.Equal(t, w.beta, 0.0)
}

func Test_WithNodes(t *testing.T) {

	w1 := NewGraph(0, 0.0).WithNodes(3)
	assert.ElementsMatch(t, w1.Nodes(), []graph.Node{graph.IntNode(0), graph.IntNode(1), graph.IntNode(2)})
	assert.Len(t, w1.nodes, 3)
	assert.Equal(t, w1.n, 3)

	w2 := NewGraph(0, 0.0).WithNodes(6)
	assert.ElementsMatch(t, w2.Nodes(), []graph.Node{graph.IntNode(0),
		graph.IntNode(1),
		graph.IntNode(2),
		graph.IntNode(3),
		graph.IntNode(4),
		graph.IntNode(5)})
	assert.Len(t, w2.nodes, 6)
	assert.Equal(t, w2.n, 6)
}

func Test_AddRemoveEdges(t *testing.T) {
	w := NewGraph(0, 0.0).WithNodes(3)
	w.addEdge(0, 1)
	w.addEdge(0, 1)
	w.addEdge(1, 2)

	assert.Len(t, w.Edges(), 4)

	w.addEdge(2, 0)
	assert.Len(t, w.Edges(), 6)

	w.removeEdge(2, 0)
	assert.Len(t, w.Edges(), 4)
	w.removeEdge(1, 0)
	assert.Len(t, w.Edges(), 2)
	w.removeEdge(2, 1)
	assert.Len(t, w.Edges(), 0)
}

func Test_WithShortEdges(t *testing.T) {

	w1 := NewGraph(1, 0.0).WithSeed(42).WithNodes(5).WithShortEdges()
	assert.Len(t, w1.Edges(), 10)

	w2 := NewGraph(2, 0.0).WithSeed(42).WithNodes(5).WithShortEdges()
	assert.Len(t, w2.Edges(), 20)

	w3 := NewGraph(3, 0.0).WithSeed(42).WithNodes(10).WithShortEdges()
	assert.Len(t, w3.Edges(), 60)

	w4 := NewGraph(4, 0.0).WithSeed(42).WithNodes(10).WithShortEdges()
	assert.Len(t, w4.Edges(), 80)
}

func Test_WithDistantEdges(t *testing.T) {

	w1 := NewGraph(1, 0.0).WithSeed(42).WithNodes(5).WithShortEdges().WithDistantEdges()
	assert.Len(t, w1.Edges(), 10)

	w2 := NewGraph(1, 1.0).WithSeed(42).WithNodes(5).WithShortEdges().WithDistantEdges()
	assert.Len(t, w2.Edges(), 10)

	w3 := NewGraph(1, 0.5).WithSeed(42).WithNodes(5).WithShortEdges().WithDistantEdges()
	assert.Len(t, w3.Edges(), 10)
}

func Test_betaZero(t *testing.T) {

	w1 := NewGraph(2, 0.0).WithSeed(42).WithNodes(7).WithShortEdges()
	assert.Equal(t, w1.n, 7)
	assert.Len(t, w1.Edges(), 28)

	w2 := NewGraph(2, 0.0).WithSeed(42).WithNodes(7).WithShortEdges().WithDistantEdges()
	assert.Equal(t, len(w1.Edges()), len(w2.Edges()))
}

func Test_betaOne(t *testing.T) {

	w1 := NewGraph(2, 1.0).WithSeed(42).WithNodes(10).WithShortEdges()
	w2 := NewGraph(2, 1.0).WithSeed(42).WithNodes(10).WithShortEdges().WithDistantEdges()

	assert.Equal(t, len(w1.Edges()), len(w2.Edges()))
	assert.Equal(t, 40, len(w1.Edges()))
}

func Test_betaHalf(t *testing.T) {

	w1 := NewGraph(2, 0.5).WithSeed(42).WithNodes(100).WithShortEdges()
	w2 := NewGraph(2, 0.5).WithSeed(42).WithNodes(100).WithShortEdges().WithDistantEdges()

	assert.Equal(t, len(w1.Edges()), len(w2.Edges()))
	assert.Equal(t, 400, len(w1.Edges()))

	assert.NotEqual(t, w1.Edges(), w2.Edges())
}

func Test_Valid(t *testing.T) {

	w1 := NewGraph(2, 0.5).WithSeed(42).WithNodes(100).WithShortEdges()
	w2 := NewGraph(24, 0.5).WithSeed(42).WithNodes(100).WithShortEdges()
	w3 := NewGraph(24, 0.5).WithSeed(42).WithNodes(1200).WithShortEdges()

	assert.False(t, w1.Valid()) // k not >> ln(n)
	assert.False(t, w2.Valid()) // n not >> k
	assert.True(t, w3.Valid())
}
