package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
)

func main() {
	writeTextfile("./test.tmp.txt")
	readBytesFromFile("./test.tmp.txt")

	createBinaryFile("./test.tmp.bin")
	readBytesFromFile("./test.tmp.bin")
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func writeTextfile(path string) {
	file, errCreate := os.Create(path)
	must(errCreate)

	_, errWrite := file.WriteString("Test 123")
	must(errWrite)

	errClose := file.Close()
	must(errClose)
}

func readBytesFromFile(path string) {
	data, errRead := os.ReadFile(path)
	must(errRead)

	fmt.Println(data)
}

func createBinaryFile(path string) {
	file, errCreate := os.Create(path)
	must(errCreate)
	defer file.Close()

	var _buf bytes.Buffer
	binary.Write(&_buf, binary.BigEndian, []byte{1, 2, 3})

	_, errWrite := file.Write(_buf.Bytes())
	must(errWrite)
}
