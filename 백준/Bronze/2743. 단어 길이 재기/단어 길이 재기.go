package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n string
	fmt.Fscanf(reader, "%s", &n)
	fmt.Println(len(n))
}