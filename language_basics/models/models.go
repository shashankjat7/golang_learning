package models

import "fmt"

type User struct {
	FirstName  string
	MiddleName string
	LastName   string
	Age        int32
	Gender     string
}

func (user User) OutputUserDetails() {
	fmt.Println(user.FirstName, user.LastName, user.Age, user.Gender)
}
