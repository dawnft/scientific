package matrix

import "github.com/lgutschmidt/scientific/structures/vector"

type Matrix []vector.Vector

func NewMatrix(rows, columns int) Matrix {
	mat := make(Matrix, rows)

	for i := range mat {
		mat[i] = make(vector.Vector, columns)
	}

	return mat
}

func (mat Matrix) Random() {
	for i := range mat {
		mat[i].Random()
	}
}

func (mat Matrix) Fill(value float64) {
	for i := range mat {
		mat[i].Fill(value)
	}
}

func (mat Matrix) Apply(function func(int, int, float64) float64) {
	for r := range mat {
		for c := range mat[r] {
			mat[r][c] = function(r, c, mat[r][c])
		}
	}
}
