package main

func esPar(number int) bool {
	return number%2 == 0
}

func login(username, pass string) bool {
	usernameSaved := "Juan"
	passSaved := "hola1234"

	return usernameSaved == username && passSaved == pass
}

func main() {
	valor1 := 1
	valor2 := 2

	if valor1 == 1 {
		println("Valor1 es 1")
	} else {
		println("Valor1 no es 1")
	}

	if valor1 == 1 && valor2 == 2 {
		println("Valor1 es 1 y Valor2 es 2")
	}

	if valor1 == 0 || valor2 == 2 {
		println("Uno de los 2 valores es verdad")
	}

	println(esPar(1))
	println(esPar(2))
	println(login("Juan", "hola1234"))
	println(login("Dario", "Hola1234"))

	modulo := 4 % 2

	switch modulo {
	case 0:
		println("Es par")
	default:
		println("Es impar")
	}

	switch modulo2 := 4 % 2; modulo2 {
	case 0:
		println("Es par")
	default:
		println("Es impar")
	}

	valor := 50
	switch {
	case valor > 100:
		println("Es mayor que 100")
	case valor < 0:
		println("Es menor que 0")
	default:
		println("Es otro numero")
	}
}
