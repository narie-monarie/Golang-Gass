package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now().Date())
	month := time.Now().Month().String()
	println(month)
}
