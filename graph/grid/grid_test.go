package grid

import (
	"testing"

	"futurae.com/smallworlds/graph"
	"github.com/stretchr/testify/assert"
)

func Test_WithAllNodes(t *testing.T) {
	w := NewGraph(2, 2).WithAllNodes()

	assert.ElementsMatch(t, w.nodes.slice(), []Position{
		at(0, 0),
		at(0, 1),
		at(1, 0),
		at(1, 1),
	})
}

func Test_WithShortEdges(t *testing.T) {
	w := NewGraph(2, 2).WithAllNodes().WithShortEdges(1)

	assert.ElementsMatch(t, w.edges.slice(), [][]Position{
		[]Position{at(0, 0), at(0, 1)},
		[]Position{at(0, 0), at(1, 0)},

		[]Position{at(1, 0), at(0, 0)},
		[]Position{at(1, 0), at(1, 1)},

		[]Position{at(0, 1), at(1, 1)},
		[]Position{at(0, 1), at(0, 0)},

		[]Position{at(1, 1), at(1, 0)},
		[]Position{at(1, 1), at(0, 1)},
	})
}

func Test_D3Object(t *testing.T) {
	o := graph.D3Json(NewGraph(2, 2).WithAllNodes().WithShortEdges(1))

	assert.NotNil(t, o["nodes"])
	assert.NotNil(t, o["links"])

	assert.Contains(t, o["nodes"], map[string]interface{}{
		"id":    "(0,0)",
		"group": 1,
	})

	assert.Contains(t, o["links"], map[string]interface{}{
		"source": "(0,0)",
		"target": "(1,0)",
		"value":  1,
	})
}

func Test_normalizingConst0(t *testing.T) {
	q := NewGraph(4, 4).
		WithAllNodes().
		WithShortEdges(1)

	assert.Equal(t, 15.0, q.normalizingConstFor(at(0, 0), 0))
}

func Test_normalizingConst2(t *testing.T) {
	q := NewGraph(4, 4).
		WithAllNodes().
		WithShortEdges(1)

	assert.InEpsilon(t, 3.4897, q.normalizingConstFor(at(0, 0), 2), 0.001)
}

func Test_WithDistantEdges(t *testing.T) {
	q := NewGraph(4, 4).
		WithAllNodes().
		WithShortEdges(1)

	p := NewGraph(4, 4).
		WithAllNodes().
		WithShortEdges(1).
		WithDistantEdges(1, 0)

	assert.Len(t, q.edges.slice(), 48)
	assert.Len(t, p.edges.slice(), len(q.edges.slice())+16)
}

func Test_WithDropout1(t *testing.T) {
	q := NewGraph(4, 4).
		WithAllNodes().
		WithShortEdges(1).
		WithDropout(1.0)

	assert.Len(t, q.edges.slice(), 0)
}

func Test_10000_Nodes(t *testing.T) {
	q := NewGraph(10, 10).
		WithSeed(42).
		WithAllNodes().
		WithShortEdges(1).
		WithDistantEdges(1, 1)

	assert.Len(t, q.edges.slice(), 460)
}
