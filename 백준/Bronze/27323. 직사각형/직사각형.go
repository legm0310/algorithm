package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var width, height int
	fmt.Fscanf(reader, "%d\n %d\n", &width, &height)
	fmt.Println(width * height)
}
