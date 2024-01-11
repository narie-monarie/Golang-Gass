package main

import (
	"fmt"
	"time"
)

func main() {
	go hello()
	time.Sleep(1 * time.Second)
	goodbye()
}

func hello() {
	fmt.Println("Hello World")
}

func goodbye() {
	fmt.Println("GoodBye World")
}
