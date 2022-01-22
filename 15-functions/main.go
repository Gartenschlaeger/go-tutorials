package main

import (
	"fmt"
	"os"
)

func init() {
	fmt.Println("init()")
}

func main() {
	fmt.Println("main()")

	func() {
		fmt.Println("Anonymous function in main")
	}()

	greeter()
	greeterByName(os.Getenv("USER"))

	fmt.Println(Add(10, 12))
	fmt.Println(Sum(1, 2, 4, 8, 16, 32))
	fmt.Println(GetMinMax(5, 9, 7, 3, 1, 6))
}

func greeter() {
	fmt.Println("Namstey from golang")
}

func greeterByName(name string) {
	fmt.Println("Hello", name)
}

func Add(a int, b int) int {
	return a + b
}

func Sum(values ...int) int {
	sum := 0
	for i := 0; i < len(values); i++ {
		sum += values[i]
	}

	return sum
}

func GetMinMax(values ...int) (int, int) {
	var min int
	var hasMin = false

	var max int
	var hasMax = false

	for _, val := range values {
		if !hasMin || val < min {
			hasMin = true
			min = val
		}
		if !hasMax || val > max {
			hasMax = true
			max = val
		}
	}

	return min, max
}
