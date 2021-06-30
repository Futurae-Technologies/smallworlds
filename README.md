# smallworlds

_smallworlds_ is a set of Golang modules to support building agent-based simulations.

The goal is to let the programmers focus on the semantics of simulations (e.g. deciding the
types of agents their behaviours etc) rather than focusing on constructing the worlds on top of which to
run the simulations.

Common use of these packages is as follows. A programmer constructs a particular small-world or random
graph, builds a world on top of this graph by adding context to the graph's nodes, and then defines a set
of agents that will walk this world.

The simulation consists of specifying the actual contexts as well as how agents behave in the world.
For example, we may specify that agents have home and work addresses in the world. That they prefer
to move between these nodes with certain transitions, but that they also occasionally wonder off and explore
other nodes. Assigning transition probabilities is what the programmer needs to do to run simulations.

Finally, we decided to separate the graph construction from the world construction as we see these two concepts
loosly-coupled. In particular, worlds can be constructed on top of any graph, while some projects may
prefer to only construct small-world graphs and have their own world implementations.

## Examples

Please see **examples.go** that illustrate the summary above.

Running simulations is outside the scope of _smallworlds_. A commmon pattern would be as below:
```
// create a small world
w := world.NewWorld(grid.NewGraph(100, 100).
		WithAllNodes().
		WithShortEdges(2).
		WithDistantEdges(2, 3).
		WithDropout(0.2))

// populate context
for _, n := range w.Nodes() {
    w.AddContext(n, world.Context{
		"rain":  rand.Float64(),
		"sunny": rand.Float64(),
	})
}

// create a random agent
randomAgent := world.NewAgent(w).
		WithState(w.Nodes()[0]).
		WithExploreProb(1.0) // Always explores randomly

// create a mostly stationary agent between two nodes.
human := world.NewAgent(w).
		WithState(w.Nodes()[0]).
		WithAddress(w.Nodes()[0]).
		WithAddress(w.Nodes()[1]).
		WithVisitDistribution([][]float64{
			[]float64{0.8, 0.2},
			[]float64{0.2, 0.8},
		}).
		WithK(10). // Agent chooses between 10 shortest walks when visiting an address
		WithExploreProb(0.05) // very very rarely will the agent wander off and visit a node that is not his address.

// simulation loop
for i:=0; i<100; i++ {
    randomAgent.VisitAddressOrExplore()
    human.VisitAddressOrExplore()
}

analyze(randomAgent.History)
analyze(human.History)

```

## Packages

#### graph
Core graph interfaces that constructed graph satisfy.

#### graph/random
Constructs random graphs.

#### graph/grid
Constructs small-world graphs based on the grid algorithm.

#### graph/ring
Constructs small-world graphs based on the ring algorithm.

#### world
Constructs worlds out of graphs, and specifies agents' behaviours.

#### stats
Utility functions for creating transition matrices.
