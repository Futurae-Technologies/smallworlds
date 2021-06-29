package grid

type positions map[int]map[int]struct{} // x->y->exists

func positionsFrom(p Position, maxDistance int) positions {
	acc := make(positions)

	for i := p.x - maxDistance; i < p.x+maxDistance+1; i++ {
		for j := p.y - maxDistance; j < p.y+maxDistance+1; j++ {
			if p.distance(at(i, j)) <= maxDistance {
				acc.add(at(i, j))
			}
		}
	}

	return acc
}

func (ps positions) distancesFrom(from Position) []int {
	distances := make([]int, 0, 0)

	for _, to := range ps.slice() {
		if !from.equal(to) {
			distances = append(distances, from.distance(to))
		}
	}

	return distances
}
func (ps positions) add(p Position) {
	_, ok := ps[p.x]
	if !ok {
		ps[p.x] = make(map[int]struct{})
	}
	ps[p.x][p.y] = struct{}{}
}

func (ps positions) remove(p Position) {
	if ps.has(p) {
		delete(ps[p.x], p.y)
	}
}

func (ps positions) has(p Position) bool {
	_, hasx := ps[p.x]
	if hasx {
		_, hasy := ps[p.x][p.y]
		if hasy {
			return true
		}
	}
	return false
}

func (ps positions) merge(qs positions) {
	for i := range qs {
		for j := range qs[i] {
			ps.add(at(i, j))
		}
	}
}

func (ps positions) slice() []Position {
	slice := make([]Position, 0, 0)

	for x := range ps {
		for y := range ps[x] {
			slice = append(slice, at(x, y))
		}
	}
	return slice
}
