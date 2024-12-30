package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	var a, b int
	fmt.Fscanf(r, "%d %d", &a, &b)
	fmt.Println(a * b)
}