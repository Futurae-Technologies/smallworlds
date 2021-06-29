package world

import (
	"testing"

	"futurae.com/smallworlds/graph/grid"
	"futurae.com/smallworlds/graph/ring"
)

func benchmarkRing(i int, b *testing.B) {

	for n := 0; n < b.N; n++ {
		g := ring.NewGraph(4, 0.0).WithSeed(42).WithNodes(i).WithShortEdges()
		m := NewWorld(g)
		a := NewAgent(m).WithSeed(42).
			WithAddress(g.Nodes()[0]).
			WithAddress(g.Nodes()[1]).
			WithVisitDistribution([][]float64{{0.3, 0.7}, {0.6, 0.4}}).
			WithState(g.Nodes()[0])

		a.VisitAddressOrExplore()
		a.VisitAddressOrExplore()
		a.VisitAddressOrExplore()
		a.VisitAddressOrExplore()
		a.VisitAddressOrExplore()
	}
}

func BenchmarkRing20(b *testing.B)  { benchmarkRing(20, b) }
func BenchmarkRing40(b *testing.B)  { benchmarkRing(40, b) }
func BenchmarkRing60(b *testing.B)  { benchmarkRing(60, b) }
func BenchmarkRing80(b *testing.B)  { benchmarkRing(80, b) }
func BenchmarkRing100(b *testing.B) { benchmarkRing(100, b) }

func benchmarkGrid(i int, b *testing.B) {

	for n := 0; n < b.N; n++ {
		g := grid.NewGraph(i, i).WithSeed(42).WithAllNodes().WithShortEdges(1)
		nodes := g.Nodes()

		m := NewWorld(g)
		a := NewAgent(m).WithSeed(42).
			WithAddress(nodes[0]).
			WithAddress(nodes[99]).
			WithVisitDistribution([][]float64{{0.3, 0.7}, {0.6, 0.4}}).
			WithState(nodes[0])

		a.VisitAddressOrExplore()
		a.VisitAddressOrExplore()
		a.VisitAddressOrExplore()
		a.VisitAddressOrExplore()
		a.VisitAddressOrExplore()
	}
}

func BenchmarkGrid20(b *testing.B)  { benchmarkGrid(20, b) }
func BenchmarkGrid40(b *testing.B)  { benchmarkGrid(40, b) }
func BenchmarkGrid60(b *testing.B)  { benchmarkGrid(60, b) }
func BenchmarkGrid80(b *testing.B)  { benchmarkGrid(80, b) }
func BenchmarkGrid100(b *testing.B) { benchmarkGrid(100, b) }
