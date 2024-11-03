package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var up, down, v int
	fmt.Fscanf(reader, "%d %d %d", &up, &down, &v)
	days := (v - down) / (up - down)
	if (v-down)%(up-down) != 0 {
		days++
	}
	fmt.Println(days)
}
