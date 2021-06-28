package grid

import (
	"fmt"
	"math"

	"futurae.com/smallworlds/graph"
)

type Position struct {
	x int
	y int
}

func at(x, y int) Position {
	return Position{x, y}
}

func (p Position) equal(q Position) bool {
	return (p.x == q.x) && (p.y == q.y)
}

func (p Position) within(boundX, boundY int) bool {
	return valid(p.x, boundX) && valid(p.y, boundY)
}

func valid(pos, bound int) bool {
	return (pos >= 0) && (pos < bound)
}

func (from Position) Distance(to graph.Node) int {
	return from.distance(to.(Position))
}

func (from Position) distance(to Position) int {
	return int(
		math.Abs(float64(to.x)-float64(from.x)) +
			math.Abs(float64(to.y)-float64(from.y)))
}

func (p Position) String() string {
	return fmt.Sprintf("(%d,%d)", p.x, p.y)
}
