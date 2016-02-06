package tesseract

import "github.com/lgutschmidt/scientific/structures/cube"

type Tesseract []cube.Cube

func NewTesseract(size1, size2, size3, size4 int) Tesseract {
	tes := make(Tesseract, size1)

	for i := range tes {
		tes[i] = cube.NewCube(size2, size3, size4)
	}

	return tes
}

func (tes Tesseract) Random() {
	for i := range tes {
		tes[i].Random()
	}
}

func (tes Tesseract) Fill(value float64) {
	for i := range tes {
		tes[i].Fill(value)
	}
}

func (tes Tesseract) Apply(function func(int, int, int, int, float64) float64) {
	for i := range tes {
		for j := range tes[i] {
			for k := range tes[i][j] {
				for l := range tes[i][j][k] {
					tes[i][j][k][l] = function(i, j, k, l, tes[i][j][k][l])
				}
			}
		}
	}
}
