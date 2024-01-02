package main

import (
	"fmt"
	"math"
)

type vertex struct {
	lat, long float64
}

var m map[string]vertex

func main() {
	m = make(map[string]vertex)
	m["Bell Labs"] = vertex{40.22, -74.35}
	fmt.Println(m)

	m := make(map[string]int)

	m["answer"] = 22

	_, ok := m["answer"]
	fmt.Println(ok)
	delete(m, "answer")
	fmt.Println(m)

	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(hypot(5, 12))
}
