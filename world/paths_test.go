package world

import (
	"testing"

	"futurae.com/smallworlds/graph/grid"
	"futurae.com/smallworlds/graph/ring"
	"github.com/stretchr/testify/assert"
)

func Test_kShortestPaths_Ring_4Nodes_Distant(t *testing.T) {
	m := NewWorld(ring.NewGraph(1, 0).WithSeed(42).WithNodes(0).WithShortEdges())
	m.array = [][]int{
		{0, 1, 1, 1},
		{1, 0, 1, 1},
		{1, 1, 0, 1},
		{1, 1, 1, 0}}

	kExpected := paths{path{0, 2}, path{0, 3, 2}}

	assert.EqualValues(t, m.kShortestPaths(2, 0, 2), kExpected)
}

func Test_kShortestPaths_10x10Grid(t *testing.T) {
	m := NewWorld(grid.NewGraph(10, 10).WithAllNodes().WithShortEdges(2))
	bottom := m.toInt["(0,0)"]
	top := m.toInt["(9,9)"]

	assert.Len(t, m.kShortestPaths(5, bottom, top), 5)
}

func Test_kShortestPaths_ring_0to0(t *testing.T) {
	m := NewWorld(ring.NewGraph(1, 0).WithSeed(42).WithNodes(0).WithShortEdges())
	m.array = [][]int{
		{0, 1, 1, 1},
		{1, 0, 1, 1},
		{1, 1, 0, 1},
		{1, 1, 1, 0}}

	kExpected := paths{path{0}}
	assert.EqualValues(t, m.kShortestPaths(1, 0, 0), kExpected)
}

func Test_kShortestPaths_grid_0to0(t *testing.T) {
	m := NewWorld(grid.NewGraph(10, 10).WithAllNodes().WithShortEdges(2))
	bottom := m.toInt["(0,0)"]
	top := m.toInt["(0,0)"]

	kExpected := paths{path{m.toInt["(0,0)"]}}
	assert.EqualValues(t, m.kShortestPaths(1, bottom, top), kExpected)
}

func Test_kShortestPaths_k(t *testing.T) {
	m := NewWorld(ring.NewGraph(1, 0).WithSeed(42).WithNodes(2).WithShortEdges())

	kExpected := paths{path{1, 0}, path{1, 0, 1, 0}}
	assert.EqualValues(t, kExpected, m.kShortestPaths(2, 1, 0))
}

func Test_path_copyAndAdd(t *testing.T) {
	p := newPath(0)
	q := p.copyAndAdd(1)

	assert.Equal(t, path([]int{0, 1}), q)
	assert.Equal(t, 2, q.cost(), 2)
	assert.Equal(t, 1, p.cost(), 1)
	assert.Equal(t, 1, q.end(), 1)
	assert.True(t, q.equals(q))
}

func Test_paths_contains(t *testing.T) {
	p := newPath(0)
	q := p.copyAndAdd(1)
	r := q.copyAndAdd(2)

	ps := make(paths, 0, 0)
	ps = ps.add(p)
	ps = ps.add(p)
	ps = ps.add(q)
	ps = ps.add(q)

	assert.Len(t, ps, 2)
	assert.True(t, ps.contains(q))
	assert.False(t, ps.contains(r))
}

func Test_RandomPath_ring(t *testing.T) {

	m := NewWorld(ring.NewGraph(1, 0).WithSeed(42).WithNodes(0).WithShortEdges())
	m.array = [][]int{
		{0, 1, 1, 1},
		{1, 0, 1, 1},
		{1, 1, 0, 1},
		{1, 1, 1, 0}}

	assert.Equal(t, len(m.randomPath(3, 0)), 3)
	assert.Equal(t, len(m.randomPath(4, 0)), 4)
}

func Test_RandomPath_grid(t *testing.T) {

	m := NewWorld(grid.NewGraph(10, 10).WithAllNodes().WithShortEdges(2))

	assert.Equal(t, len(m.randomPath(3, m.toInt["(0,0)"])), 3)
	assert.Equal(t, len(m.randomPath(4, m.toInt["(0,0)"])), 4)
}
