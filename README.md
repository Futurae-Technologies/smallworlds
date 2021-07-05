# smallworlds

_smallworlds_ is a set of Go modules for building graph-based agent simulations. We provide APIs
for generating small-world graphs, attaching contexts to nodes, and defining agents that traverse 
these worlds. 

Programmers thus focus on the semantics of simulations (e.g. agent transitions,
context definitions, interactions, etc). For example, we may specify that agents have home and work addresses in the world. 
That they prefer to move between these nodes with certain transitions, but that they also occasionally wonder off and explore
other nodes.

In short, we see this project as providing building blocks for running graph-based simulations.

## Running a simulation

We illustrate below how to construct a simple world with two agents, and run a simulation.
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

// create a random agent. This agent always wanders around.
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

// the output of every simulation are agents' histories (traces)
// which are then analyzed
analyze(randomAgent.History)
analyze(human.History)

```

Please see **examples.go** for further code samples.

Note that we decided to separate the graph construction from the world construction, as we see these two concepts
loosly-coupled. That is, a particular (contextual) world can be constructed over any graph (not just random and small-world)
thus some projects may prefer to inject their graph constructs. Similarly, some projects may prefer to construct their own
worlds on top of the small-world or random graphs.

## Packages

#### graph
Core graph interfaces for injecting generated graphs into worlds.

#### graph/random
Constructs random graphs.

#### graph/grid
Constructs small-world graphs based on the grid algorithm.

#### graph/ring
Constructs small-world graphs based on the ring algorithm.

#### world
Constructs worlds by attaching contexts to graphs, and constructs
probabilistic agent walks over the constructed worlds.

#### stats
Utility functions for creating transition matrices.
