package main

import "fmt"

func normalFunction() {
	println("normalFunction")
}

func argument(text string) {
	println("text")
}

func arguments(a, b int, c string) {
	fmt.Printf("%d mas %d es igual a %s\n", a, b, c)
}

func returned(number int) int {
	return number * 2
}

func multiReturn(a int) (b, c, d int) {
	return a, a + 1, a + 2
}

func main() {
	normalFunction()
	argument("hola mundo")
	arguments(5, 5, "10")
	value := returned(5)
	fmt.Println("Value:", value)

	a, b, _ := multiReturn(5)
	fmt.Println("a:", a)
	fmt.Println("b:", b)
	// fmt.Println("c:", c)

}
