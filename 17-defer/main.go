package main

import "fmt"

func main() {
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer myDefer()
	fmt.Println("main()")

}

func myDefer() {
	defer fmt.Print("\n")
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}
