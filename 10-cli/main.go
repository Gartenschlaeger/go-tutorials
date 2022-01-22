package main

import "fmt"

func main() {

	// output

	fmt.Println("Hello World!")
	fmt.Printf("Hello %v!\n", "Max")

	// input

	var age int
	fmt.Scan(&age)
	fmt.Printf("age = %v\n", age)

	var text string
	n, err := fmt.Scan(&text)
	fmt.Printf("Readed bytes %v, err %v\n", n, err)
	fmt.Printf("text = '%v'\n", text)

}
