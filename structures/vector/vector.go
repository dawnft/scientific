package vector

import "math/rand"

type Vector []float64

func NewVector(size int) Vector {
	return make(Vector, size)
}

func (vec Vector) Random() {
	vec.Apply(func(int, float64) float64 {
		return rand.Float64()
	})
}

func (vec Vector) Fill(value float64) {
	vec.Apply(func(int, float64) float64 {
		return value
	})
}

func (vec Vector) Apply(function func(int, float64) float64) {
	for i := range vec {
		vec[i] = function(i, vec[i])
	}
}

func (vec Vector) DotSum(other Vector) float64 {
	sum := 0.0
	for i := range vec {
		sum += vec[i] * other[i]
	}
	return sum
}
