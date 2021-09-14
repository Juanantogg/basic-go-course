package main

import (
	mypackage "curso_golang_platzi/src/9_accessModifiers/myPackage"
	"fmt"
)

func main() {
	car := mypackage.Car{Model: "Mustang", Year: 2021}
	fmt.Println(car)
	mypackage.PrintMessage()
}
