package grid

import "futurae.com/smallworlds/graph"

type edge struct {
	from Position
	to   Position
}

func (e edge) From() graph.Node {
	return e.from
}

func (e edge) To() graph.Node {
	return e.to
}

type edges map[int]map[int]map[int]map[int]struct{}

func (es edges) add(from Position, to Position) {
	es.addDirection(from, to)
	es.addDirection(to, from)
}

func (es edges) addDirection(from Position, to Position) {
	_, fromX := es[from.x]
	if !fromX {
		es[from.x] = make(map[int]map[int]map[int]struct{})
	}
	_, fromY := es[from.x][from.y]
	if !fromY {
		es[from.x][from.y] = make(map[int]map[int]struct{})
	}
	_, toX := es[from.x][from.y][to.x]
	if !toX {
		es[from.x][from.y][to.x] = make(map[int]struct{})
	}
	es[from.x][from.y][to.x][to.y] = struct{}{}
}

func (es edges) slice() [][]Position {
	acc := make([][]Position, 0, 0)

	for fx := range es {
		for fy := range es[fx] {
			for tx := range es[fx][fy] {
				for ty := range es[fx][fy][tx] {
					acc = append(acc, []Position{at(fx, fy), at(tx, ty)})
				}
			}
		}
	}
	return acc
}

func (es edges) contains(from, to Position) bool {
	_, fromX := es[from.x]
	if !fromX {
		return false
	}
	_, fromY := es[from.x][from.y]
	if !fromY {
		return false
	}
	_, toX := es[from.x][from.y][to.x]
	if !toX {
		return false
	}
	_, ok := es[from.x][from.y][to.x][to.y]

	return ok
}

func (es edges) remove(from, to Position) {
	if es.contains(from, to) {
		delete(es[from.x][from.y][to.x], to.y)
	}
}
