package random

import (
	"testing"

	"futurae.com/smallworlds/graph"
	"github.com/stretchr/testify/assert"
)

func Test_WithNodes(t *testing.T) {
	w := NewGraph().WithNodes(3)
	assert.ElementsMatch(t, w.Nodes(), []graph.Node{graph.IntNode(0), graph.IntNode(1), graph.IntNode(2)})

	w.WithNodes(3)
	assert.ElementsMatch(t, w.Nodes(), []graph.Node{
		graph.IntNode(0),
		graph.IntNode(1),
		graph.IntNode(2),
		graph.IntNode(3),
		graph.IntNode(4),
		graph.IntNode(5),
	})
}

func Test_WithEdges(t *testing.T) {
	w := NewGraph().WithNodes(3).WithEdges(2)

	assert.ElementsMatch(t, w.Nodes(), []graph.Node{graph.IntNode(0), graph.IntNode(1), graph.IntNode(2)})
	assert.Len(t, w.Edges(), 2)
}
