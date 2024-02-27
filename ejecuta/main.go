package main

import (
	"fmt"
	"github.com/jcgr5/tareaMatriz1/operaciones/generar"
	"github.com/jcgr5/tareaMatriz1/operaciones/sumatoria"
)

func ejecutar() {
	matrix := generar.CrearMatriz()
	fmt.Println("Matriz", matrix)
	column := sumatoria.SumColumnas(matrix)
	row := sumatoria.SumFilas(matrix)
	resultado := column * row
	fmt.Println("Sumatoria Filas = ", row)
	fmt.Println("Sumatoria columnas = ", column)
	fmt.Println("Sumatoria de Filas * Sumatoria Columnas = ", resultado)
}
