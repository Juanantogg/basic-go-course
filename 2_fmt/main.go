package main

import (
	"bufio"
	"fmt"
	"math"
	"math/big"
	"os"
)

func main() {
	helloMessage := "Hello"
	wordlMessage := "World"

	// Println
	fmt.Println(helloMessage, wordlMessage)
	fmt.Println(helloMessage, wordlMessage)

	// Printf
	nombre := "Platzi"
	cursos := 500

	fmt.Printf("%s tiene mas de %d cursos\n", nombre, cursos)
	fmt.Printf("%v tiene mas de %v cursos\n", nombre, cursos)

	// Sprintf
	message := fmt.Sprintf("%s tiene mas de %d cursos", nombre, cursos)
	fmt.Println(message)

	// Tipos de dato
	fmt.Printf("helloMessage: %T\n", helloMessage)
	fmt.Printf("cursos: %T\n", cursos)

	days := 7
	fmt.Printf("%T %[1]v \n", days)

	var pi = math.Pi
	fmt.Printf("%T %[1]v \n", pi)

	var pi32 float32 = math.Pi
	fmt.Printf("%T %[1]v \n", pi32)

	third := 1.0 / 3
	fmt.Println(third)
	fmt.Printf("%T %[1]v \n", third)
	fmt.Printf("%f \n", third)
	fmt.Printf("%.2f \n", third)

	// var long int64 = 7000000000000000000000000000000000000000000
	var long = new(big.Int)
	long.SetString("7000000000000000000000000000000000000000000", 10)
	fmt.Printf("%T %[1]v \n", long)

	var numero1 int
	var numero2 int
	var input string
	fmt.Println("Ingresar numero1:")
	// fmt.Scanf("%d", &numero1)
	fmt.Scanln(&numero1)
	fmt.Println("Ingresar numero2:")
	// fmt.Scanf("%d", &numero2)
	fmt.Scanln(&numero2)

	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		input = scanner.Text()
	}

	fmt.Println("Valor de input: ", input)
}
