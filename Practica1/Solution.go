package main

import (
	"fmt"
)

/*
 * Este método lo que hace es que dado un arreglo numérico multiplica cada
 * uno de sus elementos excepto por el í-esimo elemento, y el resultado lo
 * guardo en un arreglo de números enteros
 */

func productExceptSelf(nums []int) []int {
	/* Arreglo que se va a regresar con la respuesta*/
	answer := make([]int, len(nums))
	aux := 1
	aux2 := 1
	/*Slices para cuando recorremos de izquierda a derecha por primera vez el
	  arreglo y multiplicamos en dos en dos los elementos de este.*/
	firstrun := make([]int, 0)
	/*Segunda slice para cuando recorremos de derecha a izquierda por segunda
	  vez el arrgelo multiplicando en dos en dos los elementos de este*/
	secondrun := make([]int, 0)

	/*En este for recorremos de izquierda a derecha el arreglo dado en dos en dos
	  y el reultado se va guardando en la lista */
	for i := 0; i < len(nums); i++ {
		aux *= nums[i]
		firstrun = append(firstrun, aux)
	}

	/*En este for recorremos de derecha a izquierda el arreglo dado en dos en
	dos y el reultado se va guardando en la lista*/
	for i := len(nums) - 1; i >= 0; i-- {
		aux2 *= nums[i]
		secondrun = append(secondrun, aux2)
	}

	/*Volteamos la lista para poder multiplicar bien los elementos en sus posición*/
	reverse(secondrun)

	/*Recorremos el arreglo del parametro para poder identificar en que lugar
	estan los elementos y ver por cual número se debe de multiplicar el	i-esimo
	elemento menos 1 en la primera slice con el i-esimo elemento más 1 del segundo slice */
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			answer[i] = secondrun[i+1]
		} else if i > 0 && i < len(nums)-1 {
			answer[i] = firstrun[i-1] * secondrun[i+1]
		} else {
			answer[i] = firstrun[i-1]
		}
	}

	return answer
}

/*Función que regresa la reversa de un arreglo*/
func reverse(nums []int) {
	for i := 0; i < len(nums)/2; i++ {
		nums[i], nums[len(nums)-i-1] = nums[len(nums)-i-1], nums[i]
	}
}

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(productExceptSelf(nums))
	nums2 := []int{-1, 1, 0, -3, 3}
	fmt.Println(productExceptSelf(nums2))
	nums3 := []int{-5, -2, -1, 1, 2, 7, 6}
	fmt.Println(productExceptSelf(nums3))
	nums4 := []int{0, 0, 0, 1, 2, 7, 0, 6}
	fmt.Println(productExceptSelf(nums4))
}
