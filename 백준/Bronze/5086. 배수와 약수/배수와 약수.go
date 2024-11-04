package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n1, n2 int
	for {
		fmt.Fscanf(reader, "%d %d\n", &n1, &n2)
		if n1 == 0 && n2 == 0 {
			break
		}
		if n1%n2 == 0 {
			fmt.Println("multiple")
		} else if n2%n1 == 0 {
			fmt.Println("factor")
		} else {
			fmt.Println("neither")
		}
	}
}