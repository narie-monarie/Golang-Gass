package main

import "fmt"

func main() {
	var scores = []int{90, 70, 50, 80, 60, 85}

	for _, v := range scores {
		fmt.Println(v)
	}
}
