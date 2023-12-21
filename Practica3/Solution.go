package main

import "fmt"

/*
Funcion que dado un arreglo de números los ordenada
de manera O(n), sin usar la librería sort e imprime
el arreglo ordenado.

@param un arreglo de números entre 0 y 2
@return null
*/
func sortColors(nums []int) {
	var nums_0 int
	var nums_1 int
	var nums_2 int
	var index int

	/*Cuenta cuantos números en el arreglo*/
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			nums_0++
		}
		if nums[i] == 1 {
			nums_1++
		}
		if nums[i] == 2 {
			nums_2++
		}
	}

	/*
		Dependiendo del caso del arreglo vamos a imprimir el arreglo.
		Si solo tiene un elemento por default ya esta ordenado, así solo
		imprimimos el arreglo.
		Si el arreglo solo tiene un solo tipo de números por default esta ordenado.
		Si el arreglo solo tiene los números 1 y 2, entonces ordenamos
		el arreglo solo con esos números sin contar el 0.
		Por último si esta el ordenamiento con los tres números.
	*/
	if (nums_0 == 1 && nums_1 == 0 && nums_2 == 0) || (nums_0 == 0 && nums_1 == 1 && nums_2 == 0) || (nums_0 == 0 && nums_1 == 0 && nums_2 == 1) {
		fmt.Println(nums)
	} else if nums_0 == len(nums) || nums_1 == len(nums) || nums_2 == len(nums) {
		fmt.Println(nums)
	} else if nums_0 == 0 && nums_1 != 0 && nums_2 != 0 {
		for i := 0; i < nums_1; i++ {
			nums[i] = 1
		}
		for i := nums_1; i < nums_1+nums_2; i++ {
			nums[i] = 2
		}
		fmt.Println(nums)
	} else {
		for i := 0; i < nums_0; i++ {
			nums[i] = 0
			index = i
		}
		for i := index + 1; i < index+nums_1+1; i++ {
			nums[i] = 1
		}
		for i := index + nums_1 + 1; i < index+nums_1+nums_2+1; i++ {
			nums[i] = 2
		}
		fmt.Println(nums)
	}

}

/*
Pruebas que se utilizaron para el programa
*/
func main() {
	nums := []int{1, 2, 1, 2, 2, 1, 1, 2, 2, 2, 2}
	sortColors(nums)
}
