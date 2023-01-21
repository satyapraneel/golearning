package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

var mut sync.Mutex

var signals []string

func main() {
	// go greeter("hello")
	// greeter("world")
	websiteList := []string{
		"https://google.com",
		"https://github.com",
		"https://fb.com",
		"https://go.dev",
		"https://gmail.com",
	}

	for _, website := range websiteList {
		wg.Add(1)
		go getStatusCode(website)
	}
	wg.Wait()
	fmt.Println(signals)
}

// func greeter(s string) {
// 	for i := 0; i < 6; i++ {
// 		time.Sleep(3 * time.Millisecond)
// 		fmt.Println(s)
// 	}
// }

func getStatusCode(endpoint string) {
	defer wg.Done()
	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("errr")
		return
	}
	mut.Lock()
	signals = append(signals, endpoint)
	mut.Unlock()
	fmt.Printf("%s Status code %d\n", endpoint, res.StatusCode)

}
