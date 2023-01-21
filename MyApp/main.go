package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	// performGetRequest()
	// performPostJsonRequest()
	performPostFormRequest()
}
func performGetRequest() {
	const myUrl = "http://localhost:8090/get"

	response, err := http.Get(myUrl)
	if err != nil {
		panic(response)
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		panic("Response is not 200")
	}
	content, _ := io.ReadAll(response.Body)

	var responseString strings.Builder

	if _, err = responseString.Write(content); err != nil {
		panic(err)
	}

	fmt.Println(responseString.String())

	/** fmt.Println(string(content)) */
}

func performPostJsonRequest() {
	const myUrl = "http://localhost:8090/post"
	requestBody := strings.NewReader(`
	{
	    "golang": "it works"
	}`)

	response, err := http.Post(myUrl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	var responseString strings.Builder

	content, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}
	if _, err := responseString.Write(content); err != nil {
		panic(err)
	}

	fmt.Println(responseString.String())

}

func performPostFormRequest() {
	const myUrl = "http://localhost:8090/postform"

	data := url.Values{}
	data.Add("firstname", "Satyapraneel")
	data.Add("email", "satyapraneel@gmail.com")

	response, err := http.PostForm(myUrl, data)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	var responseString strings.Builder

	content, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}
	if _, err := responseString.Write(content); err != nil {
		panic(err)
	}

	fmt.Println(responseString.String())

}
