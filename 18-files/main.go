package main

import (
	"fmt"
	"os"
)

const filename = "./test.tmp.txt"

func main() {
	createTextfile()
	readTextfile()
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func createTextfile() {
	file, errCreate := os.Create(filename)
	must(errCreate)

	_, errWrite := file.WriteString("Test 123")
	must(errWrite)

	errClose := file.Close()
	must(errClose)
}

func readTextfile() {
	data, errRead := os.ReadFile(filename)
	must(errRead)

	fmt.Println(data)
}
