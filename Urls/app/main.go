package main

import (
	"fmt"
	"net/url"
)

func main() {
	const myUrl = "https://jsonplaceholder.typicode.com:443/comments?postId=1"
	result, _ := url.Parse(myUrl)
	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Path)
	// fmt.Println(result.Port())
	// fmt.Println(result.RawQuery)
	queryParams := result.Query()
	fmt.Println(queryParams["postId"])

	partsOfUrl := &url.URL{
		Scheme:   "https",
		Host:     "jsonplaceholder.typicode.com",
		Path:     "comments",
		RawQuery: "postId=10",
	}

	fmt.Println(partsOfUrl.String())

}
