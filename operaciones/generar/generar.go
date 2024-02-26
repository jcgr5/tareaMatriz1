package generar

import (
	"fmt"
	"math/rand"
)

func CrearMatriz() [][]int {
	m, n := rand.Intn(8)+1, rand.Intn(8)+1
	fmt.Println(m, "", n)
	var vector [][]int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			vector[i][j] = rand.Intn(8) + 1
		}
	}
	return vector
}
