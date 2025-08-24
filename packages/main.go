package main

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/siddkar/golang-basics/auth"
	"github.com/siddkar/golang-basics/user"
)

func main() {
	auth.LoginWithCredentials("Abc", "xyz")
	session := auth.GetSession()

	fmt.Println("session", session)

	user := user.User{
		Email: "user@gmail.com",
		Name:  "John Doe",
	}

	fmt.Println(user.Email, user.Name)

	color.Red(user.Email)
}
