package main

import (
	"fmt"
	"sync"
	"time"
)

func say(text string) {
	fmt.Println(text)
}

func say1(text string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(text)
}

func main() {
	// 1
	say("Hola")
	go say("Mundo")
	time.Sleep(time.Second)

	// 2
	var wg sync.WaitGroup

	fmt.Println("Hola")

	wg.Add(1)
	go say1("Mundo", &wg)

	wg.Add(1)
	go func(text string) {
		defer wg.Done()
		fmt.Println(text)
	}("Adios")

	wg.Wait()
}
