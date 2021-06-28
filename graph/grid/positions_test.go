package grid

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_equal(t *testing.T) {
	assert.True(t, at(1, 1).equal(at(1, 1)))
	assert.False(t, at(1, 1).equal(at(1, 2)))
}

func Test_oneStepAdjecant(t *testing.T) {
	ps := newPositions().withFromTo(at(2, 2), 1)

	assert.Len(t, ps.slice(), 5)
	assert.ElementsMatch(t, ps.slice(), []Position{
		at(1, 2),
		at(2, 1),
		at(2, 2),
		at(2, 3),
		at(3, 2),
	})
}

func Test_adjecant(t *testing.T) {
	ps := newPositions().withFromTo(at(2, 2), 2)

	expected := []Position{
		at(1, 1),
		at(1, 2),
		at(1, 3),
		at(2, 1),
		at(2, 2),
		at(2, 3),
		at(3, 1),
		at(3, 2),
		at(3, 3),
		at(2, 4),
		at(0, 2),
		at(4, 2),
		at(2, 0),
	}

	assert.ElementsMatch(t, ps.slice(), expected)
}

func Test_remove(t *testing.T) {
	ps := newPositions().withFromTo(at(2, 2), 2)

	assert.True(t, ps.has(at(1, 1)))
	ps.remove(at(1, 1))
	assert.False(t, ps.has(at(1, 1)))
}

func Test_distance(t *testing.T) {
	assert.Equal(t, at(0, 0).distance(at(3, 3)), 6)
}
