package sumatoria

import (
	"fmt"
)

func SumFilas(m [][]int) int {
	fmt.Println("Sumando Filas")
	fmt.Println(m)
	sum := 0
	for i := range m {
		fmt.Println("__________________________________________________________")
		fmt.Println("i: ", i)
		for j := range m[i] {
			fmt.Println(m[i][j])
			sum += m[i][j]
		}
		fmt.Println("suma fila: ", i+1, " = ", sum)
		fmt.Println("__________________________________________________________")
	}

	return sum
}
func SumColumnas(m [][]int) int {
	fmt.Println("Sumando Columnas")
	fmt.Println(m)
	sum := 0
	for j := range m[0] {
		fmt.Println("__________________________________________________________")
		fmt.Println("j: ", j)
		for i := range m {
			fmt.Println(m[i][j])
			sum += m[i][j]
		}
		fmt.Println("suma columna: ", j+1, " = ", sum)
		fmt.Println("__________________________________________________________")
	}
	return sum
}
