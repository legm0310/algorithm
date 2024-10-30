package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var t int
	fmt.Fscanf(reader, "%d\n", &t)
	for i := 0; i < t; i++ {
		var c int
		var q, d, n, p float64
		fmt.Fscanf(reader, "%d\n", &c)
		q = math.Floor(float64(c) / 25)
		c = c % 25
		d = math.Floor(float64(c) / 10)
		c = c % 10
		n = math.Floor(float64(c) / 5)
		c = c % 5
		p = math.Floor(float64(c) / 1)
		fmt.Printf("%d %d %d %d\n", int(q), int(d), int(n), int(p))
	}
}