package main

import (
	"fmt"
	"github.com/jcgr5/tareaMatriz1/operaciones/generar"
)

func main() {
	//var matriz = [2][2]int{}
	//matriz = generar.CrearMatriz()
	/*for i := 0; i < len(matriz); i++ {
		for j := 0; j < len(matriz[i]); j++ {
			fmt.Printf("%v", matriz[i][j])
		}
	}*/
	fmt.Println(generar.CrearMatriz())
}
