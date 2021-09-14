package main

import "fmt"

func main() {
	fmt.Println("Hola Mundo")

	// Constantes
	const pi float64 = 3.14
	const pi2 = 3.14

	fmt.Println(pi)
	fmt.Println(pi2)

	// variables

	base := 12
	var altura int = 14
	var area int

	fmt.Println(base)
	fmt.Println(altura)
	fmt.Println(area)

	// Zero values
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	ab := 10
	ac := ab * ab

	fmt.Println(ac)

	x := 10
	y := 50

	result := x + y
	fmt.Println(result)
	result = x - y
	fmt.Println(result)
	result = x * y
	fmt.Println(result)
	result = y / x
	fmt.Println(result)
	result = y % x
	fmt.Println(result)
	x++
	fmt.Println(x)
	x--
	fmt.Println(x)

}
