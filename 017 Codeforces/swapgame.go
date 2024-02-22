package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var (
	INF1 = 0x3f3f3f3f
	INF2 = int(1e17)
	N    = 1e5 + 10
	M    = 1000*100000 + 10
	mod  = 1000000007
)

type PII struct {
	x, y int
}

func solve() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	t := 1e9
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
		if i != 0 {
			t = math.Min(float64(t), float64(a[i]))
		}
	}
	if a[0] > int(t) {
		fmt.Println("Alice")
	} else {
		fmt.Println("Bob")
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	var T int
	fmt.Scan(&T)

	for ; T > 0 && scanner.Scan(); T-- {
		solve()
	}
}
