package generar

import "math/rand"

func CrearMatriz() [2]int {
	m := rand.Intn(3)
	n := rand.Intn(3)
	vector := [2]int{m, n}
	return vector
}
