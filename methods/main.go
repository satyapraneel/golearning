package main

import "fmt"

func main() {
	a := 10
	fmt.Println(a)
	user := &User{name: "Praneel"}

	user.setUser("Satyapraneel")

	user.getUser()

}

type User struct {
	name string
}

func (u *User) setUser(name string) {
	u.name = name
	fmt.Println("setting the user:", u.name)
}

func (u User) getUser() string {
	fmt.Println("getting the user: ", u.name)
	return u.name
}
