package main

import "fmt"

type pc struct {
	ram   int
	disk  int
	brand string
}

func (myPc pc) ping() {
	fmt.Println(myPc)
	fmt.Println(myPc.brand, "Pong")
}

func (myPc *pc) duplicateRam() {
	fmt.Println(myPc)
	myPc.ram *= 2
}

func (myPc pc) String() string {
	return fmt.Sprintf("Tengo %d GB Ram, %d GB Disco y es una %s", myPc.ram, myPc.disk, myPc.brand)
}

func main() {
	a := 50
	b := &a
	c := *b

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	*b = 100
	fmt.Println(a)

	myPc := pc{
		ram:   16,
		disk:  512,
		brand: "HP",
	}

	fmt.Println(myPc)
	myPc.ping()

	fmt.Println(myPc)
	myPc.duplicateRam()

	fmt.Println(myPc)
	myPc.duplicateRam()

	fmt.Println(myPc)
}
