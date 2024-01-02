package main

import (
	"fmt"
	"strings"
)

type s struct {
	i int
	b bool
}

func main() {
	v := []s{
		{2, true},
		{3, false},
		{4, false},
		{6, true},
		{7, false},
	}

	fmt.Println(v)

	a := make([]int, 5)
	a = append(a, 12)
	fmt.Println(a)

	board := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

}
