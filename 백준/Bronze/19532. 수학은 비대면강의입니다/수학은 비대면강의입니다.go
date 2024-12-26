package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	var a, b, c, d, e, f int
	fmt.Fscanf(r, "%d %d %d %d %d %d\n", &a, &b, &c, &d, &e, &f)
	for y := -999; y <= 999; y++ {
		for x := -999; x <= 999; x++ {
			if a*x+b*y == c && d*x+e*y == f {
				fmt.Println(x, y)
			}
		}
	}
}
