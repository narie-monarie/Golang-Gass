package main

import (
	"fmt"
	"slices"
)

func main() {
	//update to golang 1.22

	ma := []int{1, 3, 4, 5}
	am := []int{1, 3, 4, 5, 6}

	v := slices.Concat(ma, am)
	fmt.Println(v)
}
