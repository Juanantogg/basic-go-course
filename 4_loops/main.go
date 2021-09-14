package main

import "fmt"

func main() {
	// For conditional
	for i := 0; i <= 10; i++ {
		fmt.Println("For conditional", i)
	}

	// For while
	counter := 0
	for counter < 10 {
		fmt.Println("For while", counter)
		counter++
	}

	// For forever
	counterForever := 0
	for {
		fmt.Println(counterForever)
	}

	// 	var j = 0

	// RUTINA:
	// 	for j < 10 {
	// 		if j == 5 {
	// 			j++
	// 			fmt.Println("-")
	// 			goto RUTINA
	// 		}
	// 		fmt.Println("j: ", j)
	// 		j++
	// 	}
}
