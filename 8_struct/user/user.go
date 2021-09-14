package user

import "time"

type User struct {
	Id           int
	Name         string
	RegisterDate time.Time
	Status       bool
}

func (this *User) UserRegistry(id int, name string, registerDate time.Time, status bool) {
	this.Id = id
	this.Name = name
	this.RegisterDate = registerDate
	this.Status = status
}
