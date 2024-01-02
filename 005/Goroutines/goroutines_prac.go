package main

import (
	"fmt"
	"time"
)

// a goroutine is a lightweight thread managed by the go runtime
func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}
func main() {
	go say("world")
	say("Hello ")
}

/*
ch <- v    // Send v to channel ch.
v := <-ch  // Receive from ch, and
           // assign value to v
*/
