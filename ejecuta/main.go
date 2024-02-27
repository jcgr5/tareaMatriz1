package main

import (
	"fmt"
	"github.com/jcgr5/tareaMatriz1/operaciones/generar"
)

func main() {
	matrix := generar.CrearMatriz()
	fmt.Println("Matriz", matrix)
	column := calculos.SumColumnas(matrix)
	row := calculos.SumFilas(matrix)
	resultado := column * row
	fmt.Println("Sumatoria Filas = ", row)
	fmt.Println("Sumatoria columnas = ", column)
	fmt.Println("Sumatoria de Filas * Sumatoria Columnas = ", resultado)
}
