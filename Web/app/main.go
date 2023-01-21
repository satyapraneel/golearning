package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	const url = "https://jsonplaceholder.typicode.com/todos/1"

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	// fmt.Printf("type %T\n", response)
	data, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}
