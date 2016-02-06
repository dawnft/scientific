package body6d

import "github.com/lgutschmidt/scientific/structures/tesseract"

type Body6D [][]tesseract.Tesseract

func NewBody6D(size1, size2, size3, size4, size5, size6 int) Body6D {
	bo6 := make(Body6D, size1)

	for i := range bo6 {
		bo6[i] = make([]tesseract.Tesseract, size2)
		for j := range bo6[i] {
			bo6[i][j] = tesseract.NewTesseract(size3, size4, size5, size6)
		}
	}

	return bo6
}

func (bo6 Body6D) Random() {
	for i := range bo6 {
		for j := range bo6[i] {
			bo6[i][j].Random()
		}
	}
}

func (bo6 Body6D) Fill(value float64) {
	for i := range bo6 {
		for j := range bo6 {
			bo6[i][j].Fill(value)
		}
	}
}

func (bo6 Body6D) Apply(function func(int, int, int, int, int, int, float64) float64) {
	for i := range bo6 {
		for j := range bo6[i] {
			for k := range bo6[i][j] {
				for l := range bo6[i][j][k] {
					for m := range bo6[i][j][k][l] {
						for n := range bo6[i][j][k][l][m] {
							bo6[i][j][k][l][m][n] = function(i, j, k, l, m, n, bo6[i][j][k][l][m][n])
						}
					}
				}
			}
		}
	}
}
