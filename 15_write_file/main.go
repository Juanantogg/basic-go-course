package main

import "io/ioutil"

func catch(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	contenido := []byte("Hola Go!")
	datos := ioutil.WriteFile("prueba.txt", contenido, 0755)
	catch(datos)
}
