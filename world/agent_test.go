package world

import (
	"testing"

	"futurae.com/smallworlds/graph"
	"futurae.com/smallworlds/graph/grid"
	"futurae.com/smallworlds/graph/ring"

	"github.com/stretchr/testify/assert"
)

func Test_NewAgent(t *testing.T) {
	m := NewWorld(ring.NewGraph(1, 0).WithSeed(42).WithNodes(2))
	a := NewAgent(m)

	assert.Equal(t, m, a.world)
	assert.EqualValues(t, []graph.Node{}, a.Addresses)
	assert.EqualValues(t, map[graph.Node]map[graph.Node]float64{}, a.visitProbs)
	assert.EqualValues(t, []Walk{}, a.History)
}

func Test_AddAddress_Ring(t *testing.T) {
	g := ring.NewGraph(1, 0).WithSeed(42).WithNodes(2)
	m := NewWorld(g)
	a := NewAgent(m)
	assert.EqualValues(t, []graph.Node{}, a.Addresses)

	a.WithAddress(g.Nodes()[1])
	assert.EqualValues(t, []graph.Node{g.Nodes()[1]}, a.Addresses)

	a.WithAddress(g.Nodes()[0])
	assert.EqualValues(t, []graph.Node{g.Nodes()[1], g.Nodes()[0]}, a.Addresses)

}

func Test_AddAddress_Grid(t *testing.T) {
	g := grid.NewGraph(10, 10).WithAllNodes().WithShortEdges(1)
	nodes := g.Nodes()

	m := NewWorld(g)
	a := NewAgent(m)
	assert.EqualValues(t, []graph.Node{}, a.Addresses)

	a.WithAddress(nodes[1])
	assert.EqualValues(t, []graph.Node{nodes[1]}, a.Addresses)

	a.WithAddress(nodes[0])
	assert.EqualValues(t, []graph.Node{nodes[1], nodes[0]}, a.Addresses)
}

func Test_WithVisitProb(t *testing.T) {
	g := ring.NewGraph(1, 0).WithSeed(42).WithNodes(2).WithShortEdges()
	m := NewWorld(g)
	a := NewAgent(m).
		WithAddress(g.Nodes()[1]).
		WithAddress(g.Nodes()[0]).
		WithVisitProb(g.Nodes()[0], g.Nodes()[1], 0.5)

	assert.EqualValues(t, 0.5, a.visitProbs[g.Nodes()[0]][g.Nodes()[1]])
}

func Test_WithVisitDistribution(t *testing.T) {
	g := ring.NewGraph(1, 0).WithSeed(42).WithNodes(2).WithShortEdges()
	m := NewWorld(g)
	a := NewAgent(m).
		WithAddress(g.Nodes()[0]).
		WithAddress(g.Nodes()[1]).
		WithVisitDistribution([][]float64{{0.3, 0.7}, {0.6, 0.4}})

	assert.EqualValues(t, 0.7, a.visitProbs[g.Nodes()[0]][g.Nodes()[1]])
	assert.EqualValues(t, 0.4, a.visitProbs[g.Nodes()[1]][g.Nodes()[1]])
}

func Test_WithState_Ring(t *testing.T) {
	g := ring.NewGraph(1, 0).WithSeed(42).WithNodes(2).WithShortEdges()
	m := NewWorld(g)
	a := NewAgent(m).WithState(g.Nodes()[0])

	assert.EqualValues(t, g.Nodes()[0], a.State)
}

func Test_WithState_Grid(t *testing.T) {
	g := grid.NewGraph(10, 10).WithAllNodes().WithShortEdges(1)
	nodes := g.Nodes()

	m := NewWorld(g)
	a := NewAgent(m).WithState(nodes[0])

	assert.EqualValues(t, nodes[0], a.State)
}

func Test_End(t *testing.T) {
	g := ring.NewGraph(1, 0).WithSeed(42).WithNodes(4).WithShortEdges()

	w := Walk{g.Nodes()[0], g.Nodes()[1], g.Nodes()[2]}
	assert.EqualValues(t, g.Nodes()[2], w.End())

	w2 := Walk{g.Nodes()[0], g.Nodes()[1], g.Nodes()[0]}
	assert.EqualValues(t, g.Nodes()[0], w2.End())

}

func Test_Visit(t *testing.T) {
	g := ring.NewGraph(1, 0).WithSeed(42).WithNodes(4).WithShortEdges()
	nodes := g.Nodes()

	m := NewWorld(g)
	a := NewAgent(m).
		WithSeed(42).
		WithState(nodes[0]).
		WithAddress(nodes[0]).
		WithAddress(nodes[1]).WithK(3)

	a.Visit(nodes[2])

	assert.EqualValues(t, []Walk{{nodes[0], nodes[3], nodes[2]}}, a.History)
}

func Test_Explore(t *testing.T) {
	g := ring.NewGraph(1, 0).WithSeed(42).WithNodes(5).WithShortEdges()
	m := NewWorld(g)
	a := NewAgent(m).WithState(g.Nodes()[0]).WithMaxExploreLen(2)

	a.Explore()

	assert.Len(t, a.History, 1)
	assert.True(t, len(a.History[0]) < 3)
}

func Test_VisitOrExplore_1(t *testing.T) {
	g := ring.NewGraph(2, 0.0).WithSeed(42).WithNodes(5).WithShortEdges()
	m := NewWorld(g)
	a := NewAgent(m).WithState(g.Nodes()[0])

	a.WithK(2).WithMaxExploreLen(2).WithExploreProb(1.0)
	a.VisitOrExplore()

	assert.Len(t, a.History, 1)
	assert.True(t, len(a.History[0]) < 3)
}

func Test_VisitOrExplore_k4(t *testing.T) {
	g := ring.NewGraph(3, 0.0).WithSeed(42).WithNodes(10).WithShortEdges()
	m := NewWorld(g)
	a := NewAgent(m).
		WithState(g.Nodes()[0]).
		WithAddress(g.Nodes()[0]).
		WithAddress(g.Nodes()[5]).
		WithVisitProb(g.Nodes()[0], g.Nodes()[5], 1.0)

	a.WithK(4).WithMaxExploreLen(2).WithExploreProb(0.0)
	a.VisitOrExplore()

	assert.Len(t, a.History, 1)
	assert.Len(t, a.History[0], 3)
	assert.Equal(t, g.Nodes()[5], a.State)
}

func Test_VisitOrExplore_k1(t *testing.T) {
	g := ring.NewGraph(5, 0.0).WithSeed(42).WithNodes(10).WithShortEdges()
	m := NewWorld(g)
	a := NewAgent(m).
		WithState(g.Nodes()[0]).
		WithAddress(g.Nodes()[0]).
		WithAddress(g.Nodes()[5]).
		WithVisitProb(g.Nodes()[0], g.Nodes()[5], 1.0)

	a.WithK(1).WithMaxExploreLen(2).WithExploreProb(0.0)
	a.VisitOrExplore()

	assert.Len(t, a.History, 1)
	assert.True(t, len(a.History[0]) < 3)
	assert.Equal(t, g.Nodes()[5], a.State)
}

func Test_VisitOrExplore_fromRandom(t *testing.T) {
	g := ring.NewGraph(5, 0.0).WithSeed(42).WithNodes(10).WithShortEdges()
	m := NewWorld(g)
	a := NewAgent(m).
		WithState(g.Nodes()[1]).
		WithAddress(g.Nodes()[0]).
		WithAddress(g.Nodes()[5]).
		WithVisitProb(g.Nodes()[0], g.Nodes()[5], 1.0)

	a.WithK(1).WithMaxExploreLen(2).WithExploreProb(1.0)
	a.VisitOrExplore()

	assert.Len(t, a.History, 1)
}
