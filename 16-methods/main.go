package main

import "fmt"

type User struct {
	Id        int
	Username  string
	Firstname string
	Lastname  string
	Email     string
	IsActive  bool
}

func (u User) GetFullname() string {
	return fmt.Sprintf("%v %v", u.Firstname, u.Lastname)
}

func (u *User) ResetEmail() {
	u.Email = "noreply@kaisnet.de"
}

func main() {

	max := User{
		Id:        1,
		Username:  "maxi",
		Firstname: "Max",
		Lastname:  "Mustermann",
		Email:     "max@kaisnet.de",
		IsActive:  true}

	fmt.Println(max.GetFullname())

	max.ResetEmail()
	fmt.Println(max)

}
