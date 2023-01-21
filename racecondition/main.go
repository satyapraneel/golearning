package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	mu := &sync.RWMutex{}
	fmt.Println("race")
	var score = []int{0}
	wg.Add(3)
	// go func(wg *sync.WaitGroup, mu *sync.RWMutex) {
	// 	mu.RLock()
	// 	defer wg.Done()
	// 	fmt.Println(score)
	// 	mu.RUnlock()
	// }(wg, mu)

	go func(wg *sync.WaitGroup, mu *sync.RWMutex) {
		fmt.Println("1")
		mu.Lock()
		defer mu.Unlock()
		defer wg.Done()
		score = append(score, 1)

	}(wg, mu)
	go func(wg *sync.WaitGroup, mu *sync.RWMutex) {
		fmt.Println("2")
		mu.Lock()
		defer mu.Unlock()
		defer wg.Done()
		score = append(score, 2)
	}(wg, mu)
	go func(wg *sync.WaitGroup, mu *sync.RWMutex) {
		fmt.Println("3")
		mu.RLock()
		defer wg.Done()
		score = append(score, 3)
		mu.RUnlock()
	}(wg, mu)

	wg.Wait()
	fmt.Println(score)

}
