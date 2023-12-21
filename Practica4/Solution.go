package main

import "fmt"

// Funcion que recibe como parametros una matriz mxn,
// donde busca la suma mas corta entre el numero en la posicion
// 0x0 hasta mxn, donde solo se puede mover a la derecha y
// a la izquierda, buscando así la suma más corta para poder llegar.
func minPathSum(grid [][]int) int {
	filas, columnas := len(grid)-1, len(grid[0])-1

	matrix := make([][]int, len(grid)) //declaración de matrices auxiliar para guardar los datos de la suma
	for i := 0; i < len(grid); i++ {
		matrix[i] = make([]int, len(grid[0]))
	}

	for i := 0; i <= filas; i++ {
		for j := 0; j <= columnas; j++ {
			if i == 0 && j != 0 {
				matrix[i][j] = grid[i][j] + matrix[i][j-1]
			} else if j == 0 && i != 0 {
				matrix[i][j] = grid[i][j] + matrix[i-1][j]
			} else if j != 0 && i != 0 {
				matrix[i][j] = grid[i][j] + min(matrix[i-1][j], matrix[i][j-1])
			}

		}
	}
	return matrix[filas][columnas] + grid[0][0]
}

// Funcion auxiliar la cual nos indica que número es menor
func min(valor1 int, valor2 int) int {
	if valor1 > valor2 {
		return valor2
	}
	return valor1
}

// Pruebas que se hizo
// (pruebas del leetcode)
func main() {
	grid := [][]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(minPathSum(grid))
	grid2 := [][]int{{1, 2}, {5, 6}, {1, 1}}
	fmt.Println(minPathSum(grid2))
	grid3 := [][]int{{1, 2}, {1, 1}}
	fmt.Println(minPathSum(grid3))
	grid4 := [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}
	fmt.Println(minPathSum(grid4))
}
