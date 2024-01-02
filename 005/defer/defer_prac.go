package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func main() {

	defer fmt.Print(add(2, 4))
	fmt.Println("The result is ")

}
