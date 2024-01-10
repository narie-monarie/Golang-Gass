package main

import (
	"fmt"
	"time"
)

type Dur struct {
}

func currState(t time.Time, amount int) []int {
	m := make(map[string]int, 6)
	m[t.Month().String()] = amount
	x := []int{}
	return append(x, m[t.Month().String()])

}

func main() {
	fmt.Println(currState(time.Now(), 7000))
}
