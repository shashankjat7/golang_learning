package main

import (
	"fmt"

	"example.com/golang-learning/models"
	"example.com/golang-learning/utils"
)

func main() {
	var newUser = models.User{
		FirstName: "Shashank",
		LastName:  "Jat",
		Age:       26,
		Gender:    "Male",
	}
	newUser.OutputUserDetails()
	/// call the function to get the capital of name
	var capitalName string = utils.GetStringCapital(newUser.FirstName)
	fmt.Println(capitalName)
}
