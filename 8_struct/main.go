package main

import (
	us "./user"
	"fmt"
)

// type user struct {
// 	Id          int
// 	Name        string
// 	RegisterDay time.Time
// 	Status      bool
// }

// Hierarchy
type pepe struct {
	us.User
}

type Car struct {
	brand string
	year  int
}

func main() {
	User := new(pepe)
	fmt.Println(1, User)
	User.Id = 1
	User.Name = "Juan"
	fmt.Println(2, User)

	myCar := Car{
		brand: "Ford",
		year:  2020,
	}
	fmt.Println(myCar)

	var otherCar Car
	otherCar.brand = "Nissan"
	fmt.Println(otherCar)
}
