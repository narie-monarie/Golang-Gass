package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go hello(&wg)
	wg.Wait() // wait until innner counter has reached zero
	goodbye()
}

func hello(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Hello World")
}

func goodbye() {
	fmt.Println("GoodBye World")
}
