package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan string)
	wg.Add(1)
	go greet(ch, &wg)
	wg.Wait()
	fmt.Println("Main is Ready")
	greeting := <-ch
	time.Sleep(2 * time.Second)
	fmt.Println(greeting)

}

func greet(ch chan string, wg *sync.WaitGroup) {
	wg.Done()
	ch <- "Hello World"
	fmt.Println("Greeting Done")
}
