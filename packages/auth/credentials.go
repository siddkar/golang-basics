package auth

import "fmt"

func LoginWithCredentials(username, password string) { // default can be used in same package
	fmt.Println("Login user using credentials", username, password)
}
