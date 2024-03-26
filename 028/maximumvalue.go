package main

import "fmt"

func main() {
	var temp int
	var scores = []int{90, 70, 50, 80, 60, 85, 101}
	for i := 0; i < len(scores)-1; i++ {
		if scores[i] > scores[i+1] {
			temp = scores[i]
			scores[i] = scores[i+1]
			scores[i+1] = temp
		}
	}
	fmt.Println(scores[len(scores)-1])
}
