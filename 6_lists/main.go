package main

import "fmt"

func esPalindromo(texto string) bool {
	var textRevers string

	for i := len(texto) - 1; i >= 0; i-- {
		textRevers += string(texto[i])
	}

	fmt.Println(texto, textRevers, texto == textRevers)

	if texto == textRevers {
		return true
	} else {
		return false
	}
}

func main() {
	// Array
	var array [4]int
	fmt.Println("array", array, len(array), cap(array))
	array[1] = 1
	array[2] = 2
	fmt.Println("array", array, len(array), cap(array))

	// Slice
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("slice", slice, len(slice), cap(slice))

	fmt.Println(slice[0])
	fmt.Println(slice[:3])
	fmt.Println(slice[2:5])
	fmt.Println(slice[5:])

	slice = append(slice, 11)
	fmt.Println("slice", slice, len(slice), cap(slice))

	newSlice := []int{12, 13, 14, 15, 16, 17, 18, 19, 20}
	slice = append(slice, newSlice...)
	fmt.Println("slice", slice, len(slice), cap(slice))
	slice = append(slice, 21)
	fmt.Println("slice", slice, len(slice), cap(slice))

	fmt.Println("'Hola' es palindromo", esPalindromo("Hola"))
	fmt.Println("'ama' es palindromo", esPalindromo("ama"))

	for i, valor := range slice {
		fmt.Println(i, valor)
	}
}
