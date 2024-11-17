package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var num int
	tmp := 2
	fmt.Fscanf(reader, "%d\n", &num)
	if num == 1 {
		return
	}
	for {
		if num == 1 {
			break
		}
		if num%tmp == 0 {
			fmt.Println(tmp)
			num = num / tmp
		} else {
			tmp++
		}
	}
}
