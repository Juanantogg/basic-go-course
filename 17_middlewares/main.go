package main

import "fmt"

func main() {
	fmt.Println("main")
	fmt.Println(middleware(suma)(1, 2))
	fmt.Println(middleware(resta)(2, 3))
	fmt.Println(middleware(multiplicacion)(3, 4))
	fmt.Println(middleware(division)(4, 5))
}

func middleware(f func(int, int) int) func(int, int) int {
	return func(a, b int) int {
		fmt.Println("middleware")

		return f(a, b)
	}
}

func suma(a, b int) int {
	return a + b
}
func resta(a, b int) int {
	return a - b
}
func multiplicacion(a, b int) int {
	return a * b
}
func division(a, b int) int {
	return a / b
}
