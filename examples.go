package smallworlds

import (
	"futurae.com/smallworlds/graph"
	"futurae.com/smallworlds/graph/grid"
	"futurae.com/smallworlds/graph/random"
	"futurae.com/smallworlds/graph/ring"
	"futurae.com/smallworlds/world"
)

/*
The examples below illustrate the most common
graph and world APIs. Their detailed descriptions
are in their respective packages.
*/

// ExampleRandomGraph creates a random graph with 100 nodes and 100 edges.
func ExampleRandomGraph() graph.Graph {
	return random.NewGraph().WithNodes(100).WithEdges(100)
}

// ExampleGridSmallWorldGraph creates a 10000 node small-world graph.
// The construction follows Kleinberg's construction with short edges
// conecting neighbours of neighbours, and each node gets 2 distant edges.
// 20% of edges are randomly dropped after the small world graph has been constructed.
func ExampleGridSmallWorldGraph() graph.Graph {
	return grid.NewGraph(100, 100).
		WithAllNodes().
		WithShortEdges(2).
		WithDistantEdges(2, 3).
		WithDropout(0.2)
}

// ExampleRingSmallWorldGraph generates a small-world graph
// based on Watts-Strogatz algorithm. The generated graph has
// 1000 nodes, with 10 short connections per node on each side of the ring.
// The beta is to 0.5.
func ExampleRingSmallWorldGraph() graph.Graph {
	return ring.NewGraph(10, 0.5).
		WithNodes(1000).
		WithShortEdges().
		WithDistantEdges()
}

// ExampleWalks sets up a world using the given graph.
// We then generate a random walk of length 10,
// Followed by 3 shortest path from the node 0 to node 10.
func ExampleWalks(g graph.Graph) []world.Walk {
	w := world.NewWorld(g)

	randomWalk := w.RandomWalk(10, g.Nodes()[0])
	shortestWalks := w.KShortestPaths(3, g.Nodes()[0], g.Nodes()[10])

	return append(shortestWalks, randomWalk)
}

// ExamplePopulateWithContext adds a simple context to a node in the world.
func ExamplePopulateWithContext(w *world.World) {
	w.AddContext(w.Nodes()[0], world.Context{
		"rain":  0.1,
		"sunny": 0.9,
	})
}

// ExampleAgentWalks generates an agent in the given world.
// The agent starts off from Node 0. It has 2 addresses that it may
// travel between)
func ExampleAgentWalks(w *world.World) []world.Walk {
	a := world.NewAgent(w).
		WithState(w.Nodes()[0]).
		WithAddress(w.Nodes()[0]).
		WithAddress(w.Nodes()[1]).
		WithVisitDistribution([][]float64{
			[]float64{0.8, 0.2}, // e.g. probability of agent moving to Node 1 from Node 0 is 0.2,
			[]float64{0.5, 0.5}, // while remaining in Node 0 is 0.8.
		}).
		WithK(3). // Agent chooses between 3 shortest walks when visiting an address
		WithExploreProb(0.35)

	a.Visit(w.Nodes()[2])     // Agent travels to Node 2.
	a.Explore()               // Agent takes a random walk from Node 2.
	a.VisitAddressOrExplore() // Agent either visits one of his addresses, or will take another random walk.

	return a.History
}
