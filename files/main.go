package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fileName := "./test.txt"
	writeToFile("This needs to be go in file", fileName)
	readFromFile(fileName)
}

func writeToFile(content, fileName string) {
	file, err := os.Create(fileName)
	checkNilError(err)
	defer file.Close()
	_, err = io.WriteString(file, content)
	checkNilError(err)
}

func readFromFile(fileName string) {
	dataBytes, err := os.ReadFile(fileName)
	checkNilError(err)
	fmt.Println(string(dataBytes))
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
