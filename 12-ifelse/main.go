package main

import "fmt"

func main() {

	count := 0

	fmt.Print("Enter a count ")
	fmt.Scan(&count)

	var result string
	if count <= 10 {
		result = "<= 10"
	} else if count > 100 {
		result = "> 100"
	} else if count > 10 {
		result = "> 10"
	} else {
		result = "else"
	}

	fmt.Println("Count", result)

	if count%2 == 0 {
		fmt.Println("Count is even")
	} else {
		fmt.Println("Count is not even")
	}

	if num := 3; num > 10 {
		fmt.Println("num > 10")
	} else {
		fmt.Println("num < 10")
	}

}
