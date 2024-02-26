package generar

import (
	"math/rand"
)

func CrearMatriz() [][]int {
	m, n := rand.Intn(8)+1, rand.Intn(8)+1

	var vector [][]int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			vector[i][j] = rand.Intn(8) + 1
		}
	}
	return vector
}
