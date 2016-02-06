package cube

import "github.com/lgutschmidt/scientific/structures/matrix"

type Cube []matrix.Matrix

func NewCube(size1, size2, size3 int) Cube {
	ten := make(Cube, size1)

	for i := range ten {
		ten[i] = matrix.NewMatrix(size2, size3)
	}

	return ten
}

func (ten Cube) Random() {
	for i := range ten {
		ten[i].Random()
	}
}

func (ten Cube) Fill(value float64) {
	for i := range ten {
		ten[i].Fill(value)
	}
}

func (ten Cube) Apply(function func(int, int, int, float64) float64) {
	for i := range ten {
		for j := range ten[i] {
			for k := range ten[i][j] {
				ten[i][j][k] = function(i, j, k, ten[i][j][k])
			}
		}
	}
}
