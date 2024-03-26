package main

import "fmt"

func main() {

	var scores = []int{90, 70, 50, 80, 60, 85, 101}

	for i := 0; i < len(scores)-1; i++ {
		for j := 0; j < len(scores)-i-1; j++ {
			if scores[j] > scores[j+1] {
				var temp = scores[j]
				scores[j] = scores[j+1]
				scores[j+1] = temp
			}
		}
	}
	fmt.Println(scores)
}
