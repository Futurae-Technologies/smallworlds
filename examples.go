package smallworlds

import (
	"futurae.com/smallworlds/graph"
	"futurae.com/smallworlds/graph/grid"
	"futurae.com/smallworlds/graph/random"
	"futurae.com/smallworlds/graph/ring"
)

func NewRandomGraph() graph.Graph {
	return random.NewGraph().WithNodes(100).WithEdges(100)
}

func NewGridSmallWorldGraph() graph.Graph {
	return grid.NewGraph(100, 100).
		WithAllNodes().
		WithShortEdges(2).
		WithDistantEdges(2, 3).
		WithDropout(0.2)
}

func NewRingSmallWorldGraph() graph.Graph {
	return ring.NewGraph(10, 0.5).
		WithNodes(1000).
		WithShortEdges().
		WithDistantEdges()
}
