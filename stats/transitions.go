package stats

import (
	"math"
	"math/rand"
)

// UniformTransitionMatrix creates a uniform transition matrix
// between n points.
func UniformTransitionMatrix(n int) [][]float64 {
	prob := make([][]float64, n, n)

	for i := 0; i < n; i++ {
		prob[i] = make([]float64, n, n)
		for j := 0; j < n; j++ {
			prob[i][j] = 1 / float64(n)
		}
	}
	return prob
}

// RandomTransitionMatrix creates a random transition matrix
// between n points.
func RandomTransitionMatrix(r *rand.Rand, n int) [][]float64 {
	prob := make([][]float64, n, n)

	for i := 0; i < n; i++ {
		prob[i] = make([]float64, n)

		for j := 0; j < n; j++ {
			prob[i][j] = r.Float64()
		}
		prob[i] = Softmax(prob[i])
	}
	return prob
}

// Softmax applies softmax function over the given array.
func Softmax(array []float64) []float64 {
	exp := make([]float64, len(array))
	expSum := 0.0

	for i := 0; i < len(array); i++ {
		exp[i] = math.Exp(array[i])
		expSum += exp[i]
	}

	for i := 0; i < len(array); i++ {
		exp[i] = exp[i] / expSum
	}

	return exp
}

// PickFromDiscreteDist selects the position of the dist array
// given the value of the position.
func PickFromDiscreteDist(dist []float64) int {
	r := rand.Float64()
	sum := 0.0
	for i, prob := range dist {
		sum += prob
		if r < sum {
			return i
		}
	}
	return len(dist) - 1
}
