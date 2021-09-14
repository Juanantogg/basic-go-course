package main

import "fmt"

func say2(text string, c chan<- string) {
	c <- text
}

func main() {
	c := make(chan string, 1)

	fmt.Println("Hello")

	go say2("Bey", c)

	fmt.Println(<-c)
}
