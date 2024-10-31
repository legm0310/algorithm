package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n, d int
	s := 1
	fmt.Fscanf(reader, "%d\n", &n)
	for i := 0; i < n; i++ {
		s *= 4
		line := int(math.Sqrt(float64(s))) + 1
		d = line * line
	}
	fmt.Println(d)
}
