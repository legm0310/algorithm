package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var cnt int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanf(reader, "%d\n", &cnt)
	if cnt < 2 {
		fmt.Printf("%d", 0)
		return
	}
	var maxX, minX, maxY, minY int
	for i := 0; i < cnt; i++ {
		var x, y int
		fmt.Fscanln(reader, &x, &y)
		if i == 0 {
			maxX = x
			minX = x
			maxY = y
			minY = y
		}
		if x > maxX {
			maxX = x
		}
		if x < minX {
			minX = x
		}
		if y > maxY {
			maxY = y
		}
		if y < minY {
			minY = y
		}
	}
	fmt.Printf("%d\n", (maxX - minX) * (maxY - minY))
}
