package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string   `json:"name"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	decodeJson()
}

func encodeJson() []byte {
	jsonContent := []course{
		{"test", "1234", []string{"xyz", "pqr"}},
		{"test1", "12344", nil},
	}

	finalContent, err := json.MarshalIndent(jsonContent, "", "  ")

	if err != nil {
		panic(err)
	}
	fmt.Println(string(finalContent))

	return finalContent
	// fmt.Printf("%s\n", finalContent)
}

func decodeJson() {
	jsonData := []byte(`{
		"name": "test",
		"password":"123",
		"tags": [
		  "xyz",
		  "pqr"
		]
	  }`)
	isJson := json.Valid(jsonData)
	if !isJson {
		fmt.Println("Not a valid json")

		return
	}
	var course course
	if err := json.Unmarshal(jsonData, &course); err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", course)

	var myData map[string]any
	json.Unmarshal(jsonData, &myData)
	fmt.Printf("%#v\n", myData)

}
