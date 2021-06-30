package world

import (
	"fmt"
	"math/rand"
	"time"

	"futurae.com/smallworlds/graph"
	"futurae.com/smallworlds/stats"
)

// Agent type defines a single agent within a world.
// An agent can have addresses (nodes of interest) and the transition matrix between them.
// That is, how likely it is for an agent to travel from one address to another.
//
// An agent explores the world (i.e. takes random walks) with _exploreProb_ probability when
// given a chance.
//
// All walks are recorded in the agent's history as well as the last node that he visited.
type Agent struct {
	Addresses   []graph.Node
	Transitions map[graph.Node]map[graph.Node]float64

	world   *World
	State   graph.Node
	History []Walk

	rand *rand.Rand

	maxExploreLen int
	exploreProb   float64
	k             int
}

// NewAgent initializes an agent in the given world.
// A new agent needs to be further defined using builder functions to
// its addresses, transitions, as well as the start state.
func NewAgent(w *World) *Agent {
	return &Agent{
		world:       w,
		Addresses:   make([]graph.Node, 0),
		Transitions: make(map[graph.Node]map[graph.Node]float64),
		History:     make([]Walk, 0, 0),
		rand:        rand.New(rand.NewSource(time.Now().UTC().UnixNano())),

		maxExploreLen: 4,
		exploreProb:   0.3,
		k:             5,
	}
}

// WithSeed is a builder that sets the random number generator.
func (a *Agent) WithSeed(seed int64) *Agent {
	a.rand = rand.New(rand.NewSource(seed))
	return a
}

// WithK sets the maximum number of shortest routes that an agent chooses
// from when traversing between its addresses.
func (a *Agent) WithK(k int) *Agent {
	a.k = k
	return a
}

// WithMaxExploreLen sets the maximum length of random walks that agent may take.
func (a *Agent) WithMaxExploreLen(i int) *Agent {
	a.maxExploreLen = i
	return a
}

// WithExploreProb sets the probability of an agent abandoning
// his move to one his addresses and instead takes a random walk.
func (a *Agent) WithExploreProb(f float64) *Agent {
	a.exploreProb = f
	return a
}

// WithAddresses adds the given addresses to the agent's addresses.
func (a *Agent) WithAddresses(as []graph.Node) *Agent {
	for _, i := range as {
		a.Addresses = append(a.Addresses, i)
		a.Transitions[i] = make(map[graph.Node]float64)
	}
	return a
}

// WithAddress adds the given address to the agent.
func (a *Agent) WithAddress(ad graph.Node) *Agent {
	a.Addresses = append(a.Addresses, ad)
	a.Transitions[ad] = make(map[graph.Node]float64)
	return a
}

// WithVisitProb assigns the probability that an agent will traverse
// to the given address when at the from address.
func (a *Agent) WithVisitProb(from, to graph.Node, p float64) *Agent {
	a.Transitions[from][to] = p
	return a
}

// WithVisitDistribution sets the transition matrix between agent's
// addresses. Note that the given distribution is n-to-n mapping
// according to the ordering of the agent's addresses.
func (a *Agent) WithVisitDistribution(d [][]float64) *Agent {
	if len(d) != len(a.Addresses) {
		panic(fmt.Sprintf("Bad distribution"))
	}

	for i := 0; i < len(d); i++ {
		for j := 0; j < len(d); j++ {
			a.WithVisitProb(a.Addresses[i], a.Addresses[j], d[i][j])
		}
	}

	return a
}

// WithState positions the agent in the world.
func (a *Agent) WithState(s graph.Node) *Agent {
	a.State = s
	return a
}

// Visit moves the agent from the current state
// to the given node. The agent picks randomly between
// k shortest paths to that location.
func (a *Agent) Visit(to graph.Node) {
	ps := a.world.KShortestPaths(a.k, a.State, to)
	w := ps[0]
	if len(ps) > 1 {
		w = ps[a.rand.Intn(len(ps)-1)]
	}
	a.State = w.End()
	a.History = append(a.History, w)
}

// Explore picks randomly the length of a random walk
// and then sets the agent on it.
func (a *Agent) Explore() {
	length := a.rand.Intn(a.maxExploreLen) + 1

	w := a.world.RandomWalk(length, a.State)
	a.State = w.End()
	a.History = append(a.History, w)
}

// VisitOrExplore decides whether to Visit (an address)
// or Explore the world. If the agent decides to visit
// an address and the agent is not one of his addresses, then
// the agent will walk to one of this addresses picked at random.
func (a *Agent) VisitAddressOrExplore() {
	if a.rand.Float64() < a.exploreProb {
		a.Explore()
		return
	}

	// If Agent is not at one of his addresses, then it visits one of them at random.
	if !graph.Contains(a.Addresses, a.State) {
		a.Visit(a.Addresses[a.rand.Intn(len(a.Addresses))])
		return
	}

	// If Agent is at one of his addresses, then use the transition matrix to visit another address.
	keys, probs := a.transitionsFrom(a.State)
	a.Visit(keys[stats.PickFromDiscreteDist(probs)]) // walk there given some short path.
}

func (a *Agent) transitionsFrom(n graph.Node) ([]graph.Node, []float64) {
	probs := make([]float64, 0, 0)
	keys := make([]graph.Node, 0, 0)

	for key, value := range a.Transitions[a.State] { // pick the distribution from the state
		probs = append(probs, value)
		keys = append(keys, key)
	}

	return keys, probs
}
