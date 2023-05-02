package model

import "fmt"

type User struct {
	User     string
	Password string
}

func (u *User) printUser() {
	fmt.Println("USer:", u.User)
}
