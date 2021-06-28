package stats

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Softmax_Equal(t *testing.T) {
	expected := []float64{0.5, 0.5}
	ts := Softmax([]float64{1, 1})

	assert.Equal(t, expected, ts)
}

func Test_Softmax_4(t *testing.T) {
	expected := []float64{0.0021, 0.0058, 0.1182, 0.8737}
	ts := Softmax([]float64{-1, 0, 3, 5})

	assert.InDeltaSlice(t, expected, ts, 0.001)
}

func Test_PickFromDiscreteDist_1(t *testing.T) {
	assert.Equal(t, 0, PickFromDiscreteDist([]float64{1}))
	assert.Equal(t, 0, PickFromDiscreteDist([]float64{0.1}))
}

func Test_PickFromDiscreteDist_2(t *testing.T) {
	assert.Equal(t, 0, PickFromDiscreteDist([]float64{1, 0}))
	assert.Equal(t, 1, PickFromDiscreteDist([]float64{0, 1}))
}

func Test_UniformTransitionMatrix(t *testing.T) {
	expected := [][]float64{
		[]float64{0.25, 0.25, 0.25, 0.25},
		[]float64{0.25, 0.25, 0.25, 0.25},
		[]float64{0.25, 0.25, 0.25, 0.25},
		[]float64{0.25, 0.25, 0.25, 0.25},
	}
	assert.Equal(t, expected, UniformTransitionMatrix(4))
}

func Test_RandomTransitionMatrix(t *testing.T) {
	m := RandomTransitionMatrix(rand.New(rand.NewSource(42)), 3)
	for i := 0; i < 3; i++ {
		sum := 0.0
		for j := 0; j < 3; j++ {
			sum = sum + m[i][j]
		}
		assert.Equal(t, 1.0, sum)
	}
}
