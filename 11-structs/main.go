package main

import "fmt"

func main() {

	max := User{id: 1, Firstname: "Max", Lastname: "Mustermann"}

	fmt.Println(max)
	fmt.Printf("%+v\n", max)
	fmt.Printf("Hello %v %v\n", max.Firstname, max.Lastname)
}

type User struct {
	id        uint
	Firstname string
	Lastname  string
}
