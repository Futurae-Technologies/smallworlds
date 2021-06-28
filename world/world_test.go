package world

import (
	"testing"

	"futurae.com/smallworlds/graph"
	"futurae.com/smallworlds/graph/grid"
	"futurae.com/smallworlds/graph/random"
	"futurae.com/smallworlds/graph/ring"
	"github.com/stretchr/testify/assert"
)

func Test_NewWorld(t *testing.T) {
	m := NewWorld(ring.NewGraph(1, 0.0).WithSeed(42).WithNodes(3).WithShortEdges())

	assert.Equal(t, m.n, 3)
	assert.Len(t, m.neighbourhood(0), 2)
	assert.Len(t, m.neighbourhood(1), 2)
	assert.Len(t, m.neighbourhood(2), 2)

	assert.Equal(t, m.array[0], []int{0, 1, 1})
	assert.Equal(t, m.array[1], []int{1, 0, 1})
	assert.Equal(t, m.array[2], []int{1, 1, 0})
}

func Test_hasEdge(t *testing.T) {
	m1 := NewWorld(ring.NewGraph(1, 0.0).WithSeed(42).WithNodes(8).WithShortEdges()) // Only local edges

	assert.Equal(t, m1.array[0][0], 0)
	assert.Equal(t, m1.array[0][1], 1)
	assert.Equal(t, m1.array[1][4], 0)
	assert.Equal(t, m1.array[3][0], 0)

	m2 := NewWorld(ring.NewGraph(1, 0.0).WithSeed(42).WithNodes(8).WithShortEdges().WithDistantEdges()) // With distant edges and beta=0
	assert.Equal(t, m2.array[0][0], 0)
	assert.Equal(t, m2.array[0][1], 1)
	assert.Equal(t, m2.array[1][4], 0)
	assert.Equal(t, m2.array[3][0], 0)
}

func Test_localClusteringCoeffFor_With4(t *testing.T) {
	m1 := NewWorld(ring.NewGraph(1, 0.0).WithSeed(42).WithNodes(4).WithShortEdges())
	assert.Len(t, m1.neighbourhood(0), 2)
	assert.InEpsilon(t, m1.localClusteringCoeffFor(0), 0.6666, 0.001)
}

func Test_localClusteringCoeffFor_With3(t *testing.T) {
	m2 := NewWorld(ring.NewGraph(1, 0.0).WithSeed(42).WithNodes(3).WithShortEdges())
	assert.Len(t, m2.neighbourhood(0), 2)
	assert.Equal(t, m2.array[0], []int{0, 1, 1})
	assert.Equal(t, m2.array[0][1], 1)
	assert.Equal(t, m2.array[1][2], 1)
	assert.Equal(t, m2.array[2][0], 1)
	assert.Equal(t, m2.localClusteringCoeffFor(0), 1.0)
}

func Test_AvgClusteringCoeff_Comparison(t *testing.T) {
	m1 := NewWorld(ring.NewGraph(100, 1.0).WithSeed(42).WithNodes(10000).WithShortEdges().WithDistantEdges())
	m2 := NewWorld(grid.NewGraph(10, 10).WithSeed(42).WithAllNodes().WithShortEdges(1))

	assert.Less(t, m1.AvgClusteringCoeff(), m2.AvgClusteringCoeff())
}

func Test_AvgClusteringCoeff_Comparison_2(t *testing.T) {
	m1 := NewWorld(random.NewGraph().WithSeed(42).WithNodes(100).WithEdges(100))
	m2 := NewWorld(grid.NewGraph(10, 10).WithSeed(42).WithAllNodes().WithShortEdges(1))

	assert.Less(t, m1.AvgClusteringCoeff(), m2.AvgClusteringCoeff())
}

func Test_AvgClusteringCoeff(t *testing.T) {

	m := NewWorld(ring.NewGraph(1, 0.0).WithSeed(42).WithNodes(3).WithShortEdges())
	assert.Equal(t, m.localClusteringCoeffFor(0), 1.0)
	assert.Equal(t, m.localClusteringCoeffFor(1), 1.0)
	assert.Equal(t, m.localClusteringCoeffFor(2), 1.0)
	assert.Equal(t, m.AvgClusteringCoeff(), 1.0)

	m1 := NewWorld(ring.NewGraph(100, 1.0).WithNodes(10000).WithShortEdges().WithDistantEdges())
	m2 := NewWorld(grid.NewGraph(10, 10).WithAllNodes().WithShortEdges(1))
	assert.Less(t, m1.AvgClusteringCoeff(), m2.AvgClusteringCoeff())

}

func Test_ToNodes(t *testing.T) {
	g := ring.NewGraph(1, 0).WithSeed(42).WithNodes(5).WithShortEdges()
	m := NewWorld(g)

	assert.EqualValues(t, []graph.Node{graph.Node(g.Nodes()[0]), graph.Node(g.Nodes()[1])}, m.toNodes(path([]int{0, 1})))
}

func Test_RandomWalk(t *testing.T) {

	g := ring.NewGraph(2, 0).WithSeed(42).WithNodes(5).WithShortEdges()
	m := NewWorld(g).WithSeed(42)

	assert.EqualValues(t, Walk{g.Nodes()[0], g.Nodes()[2]}, m.RandomWalk(2, g.Nodes()[0]))
}

func Test_KShortestPaths(t *testing.T) {
	g := ring.NewGraph(2, 0).WithSeed(42).WithNodes(5).WithShortEdges()
	m := NewWorld(g)

	w := m.KShortestPaths(2, g.Nodes()[0], g.Nodes()[1])
	assert.EqualValues(t, []Walk{
		Walk{g.Nodes()[0], g.Nodes()[1]},
		Walk{g.Nodes()[0], g.Nodes()[4], g.Nodes()[1]},
	}, w)
}

func Test_EmptyContexts(t *testing.T) {

	g := ring.NewGraph(2, 0).WithSeed(42).WithNodes(5).WithShortEdges()
	m := NewWorld(g).WithSeed(42)

	assert.EqualValues(t, Walk{g.Nodes()[0], g.Nodes()[2]}, m.RandomWalk(2, g.Nodes()[0]))
	assert.EqualValues(t, []Context{{}, {}}, m.Contexts(m.RandomWalk(2, g.Nodes()[0])))
}

func Test_ShortestPaths_2(t *testing.T) {
	g := ring.NewGraph(1, 0).WithSeed(42).WithNodes(2).WithShortEdges()
	m := NewWorld(g).WithSeed(42)
	ds := m.ShortestPathsLens()

	assert.Equal(t, 2, sumDistances(ds))
}

func Test_ShortestPaths_3(t *testing.T) {
	g := ring.NewGraph(1, 0).WithSeed(42).WithNodes(3).WithShortEdges()
	m := NewWorld(g).WithSeed(42)
	ds := m.ShortestPathsLens()

	assert.Equal(t, 6, sumDistances(ds))
}

func Test_ShortestPaths_5(t *testing.T) {
	g := ring.NewGraph(1, 0).WithSeed(42).WithNodes(5).WithShortEdges()
	m := NewWorld(g).WithSeed(42)
	ds := m.ShortestPathsLens()

	assert.Equal(t, 30, sumDistances(ds))
}

func sumDistances(m [][]int) int {
	sum := 0
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m); j++ {
			sum = sum + m[i][j]
		}
	}
	return sum
}
