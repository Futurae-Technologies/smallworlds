package world

import (
	"container/heap"
)

type path []int

func newPath(v int) path {
	return []int{v}
}

func (p path) copyAndAdd(v int) path {
	q := make(path, 0, 0)

	for i := 0; i < len(p); i++ {
		q = append(q, p[i])
	}
	return append(q, v)
}

func (p path) cost() int {
	return len(p)
}

func (p path) end() int {
	return p[len(p)-1]
}

func (p path) equals(q path) bool {
	if p.cost() != q.cost() {
		return false
	}

	for i := 0; i < len(p); i++ {
		if p[i] != q[i] {
			return false
		}
	}
	return true
}

type paths []path

func (ps paths) contains(p path) bool {
	for i := 0; i < len(ps); i++ {
		if ps[i].equals(p) {
			return true
		}
	}
	return false
}

func (ps paths) add(p path) paths {
	if ps.contains(p) {
		return ps
	}

	return append(ps, p)
}

func (m *World) randomPath(length int, from int) path {
	acc := newPath(from)
	current := from

	for i := 0; i < length-1; i++ {
		ns := m.neighbourhood(current)
		current := ns[m.rand.Intn(len(ns))]
		acc = append(acc, current)
	}
	return acc
}

type indexedPath struct {
	value path
	index int
	cost  int
}

type pathHeap []*indexedPath

func (b pathHeap) Len() int {
	return len(b)
}

func (b pathHeap) Less(i, j int) bool {
	return b[i].cost < b[j].cost
}

func (b pathHeap) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
	b[i].index = i
	b[j].index = j
}

func (b *pathHeap) Push(x interface{}) {
	n := len(*b)
	i := x.(*indexedPath)
	i.index = n
	*b = append(*b, i)
}

func (b *pathHeap) Pop() interface{} {
	old := *b
	n := len(old)

	i := old[n-1]
	old[n-1] = nil

	i.index = -1
	*b = old[0 : n-1]

	return i
}

func (b pathHeap) Empty() bool {
	return b.Len() == 0
}

func (m *World) kShortestPaths(k, src, target int) paths {

	b := make(pathHeap, 0, 0)
	paths := make(paths, 0, 0)
	count := make([]int, len(m.array), len(m.array))

	heap.Init(&b)

	for i := 0; i < len(m.array); i++ {
		count[i] = 0
	}

	heap.Push(&b, &indexedPath{value: newPath(src), cost: 1})

	for {
		p := heap.Pop(&b).(*indexedPath).value

		u := p.end()
		count[u]++

		if u == target {
			paths = paths.add(p)
		}

		if count[u] <= k {
			for _, v := range m.neighbourhood(u) {
				heap.Push(&b, &indexedPath{
					value: p.copyAndAdd(v),
					cost:  p.cost() + 1})
			}
		}

		if b.Empty() || (count[target] >= k) {
			break
		}

	}
	return paths

}
