package generar

import "math/rand"

func CrearMatriz() [2][2]int {
	m, n := rand.Intn(8)+1, rand.Intn(8)+1
	vector := [2][2]int{
		{m, n},
		{n, m},
	}
	return vector
}
