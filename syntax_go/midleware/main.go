package main

import "fmt"

/*
	Son interceptores que permiten ejecutar instrucciones comunes a varias funciones que reciben
	y devuelven los mismos tipos de variables
*/
var result int

func main() {
	fmt.Println("inicio")

	result = operacionesMidd(sumar)(2, 3)
	fmt.Println(result)

	result = operacionesMidd(restar)(10, 6)
	fmt.Println(result)

	result = operacionesMidd(multiplicar)(2, 5)
	fmt.Println(result)
}

func sumar(a, b int) int {
	return a + b
}

func restar(a, b int) int {
	return a - b
}

func multiplicar(a, b int) int {
	return a * b
}

func operacionesMidd(fun func(int, int) int) func(int, int) int {
	return func(a, b int) int { //debe coincidir con lo que pide y con lo que retorna
		fmt.Println("Inicio de operacion")
		//aqui se pueden realizar cualquier operacion
		return fun(a, b)
	}
}
