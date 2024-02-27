package generar

import (
	"fmt"
	"math/rand"
)

func CrearMatriz() [][]int {
	m, n := rand.Intn(8)+1, rand.Intn(8)+1
	fmt.Println("m: ", m, ", n: ", n)
	matrix := make([][]int, m)
	for i := range matrix {
		matrix[i] = make([]int, n)
		for j := 0; j < n; j++ {
			matrix[i][j] = rand.Intn(9) + 1
		}
	}
	return matrix
}
