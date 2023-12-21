package main

import "fmt"

/*
Funcion que recibe una matriz nxn, donde cada valor
v = grid[i][j] representa una torre de v cubos colocados
encima de la celda (i, j). Y vemos la proyección de estos
cubos sobre los planos xy, yz y zx. Calculamos el area
total de la sombra

@param una matriz nxn
@return el area total de la proyeccion de la sombra
*/
func projectionArea(grid [][]int) int {
	result := 0

	for i := 0; i < len(grid); i++ {
		max_X, max_Y := 0, 0
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] != 0 {
				result++
			}
			if grid[i][j] > max_X {
				max_X = grid[i][j]
			}
			if grid[j][i] > max_Y {
				max_Y = grid[j][i]
			}
		}
		result = result + max_Y + max_X
	}
	return result

}

/*
Pruebas que se utilizaron para el programa
*/
func main() {
	test := [][]int{{1, 4}, {1, 1}}
	result := projectionArea(test)
	fmt.Println(result)
	test2 := [][]int{{1, 2}, {3, 4}}
	result2 := projectionArea(test2)
	fmt.Println(result2)
	test3 := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	result3 := projectionArea(test3)
	fmt.Println(result3)
}
