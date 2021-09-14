package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func catch(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	data, error := ioutil.ReadFile("prueba.txt")
	catch(error)
	fmt.Println(string(data))

	archivo, err := os.Open("prueba.txt")
	catch(err)
	scanner := bufio.NewScanner(archivo)
	for scanner.Scan() {
		fmt.Printf("Linea > " + scanner.Text() + "\n")
	}
	archivo.Close()
}
