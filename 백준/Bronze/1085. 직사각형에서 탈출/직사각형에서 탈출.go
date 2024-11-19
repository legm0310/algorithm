package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var curX, curY, width, height, num int
	fmt.Fscanf(reader, "%d %d %d %d\n", &curX, &curY, &width, &height)
	num = min(min(curX, width-curX), min(curY, height-curY))
	fmt.Println(num)
}
