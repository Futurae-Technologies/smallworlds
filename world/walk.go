package world

import "futurae.com/smallworlds/graph"

// Walk type is an ordered list of graph nodes.
type Walk []graph.Node

// Reverse function reverses the walk in terms
// of the visited nodes.
func (w Walk) Reverse() Walk {
	for i, j := 0, len(w)-1; i < j; i, j = i+1, j-1 {
		w[i], w[j] = w[j], w[i]
	}
	return w
}

// End returns the last visited node.
func (w Walk) End() graph.Node {
	return w[len(w)-1]
}
