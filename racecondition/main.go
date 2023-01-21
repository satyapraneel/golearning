package main

import (
	"fmt"
)

func main() {

	fmt.Println("race")
	score := []int{0}
	go func() {
		fmt.Println("1")
		score = append(score, 1)
	}()
	go func() {
		fmt.Println("2")
		score = append(score, 2)
	}()
	go func() {
		fmt.Println("3")
	}()
}
