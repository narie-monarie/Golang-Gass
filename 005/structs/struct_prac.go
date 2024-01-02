package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	vertex := Vertex{1, 2}
	p := &vertex
	p.X = 12
	fmt.Println(vertex.X)
}
