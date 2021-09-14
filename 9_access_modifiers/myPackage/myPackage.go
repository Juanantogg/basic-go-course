package mypackage

import "fmt"

type Car struct {
	Model string
	Year  int
}

type car struct {
	model string
	year  int
}

func PrintMessage() {
	fmt.Println("Mensaje")
}
