package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	numMap := make(map[int]bool)
	sum := 0
	for i := 0; i < 3; i++ {
		var num int
		fmt.Fscanf(r, "%d\n", &num)
		sum += num
		numMap[num] = true
	}
	if sum != 180 {
		fmt.Println("Error")
		return
	}
	if len(numMap) == 1 {
		fmt.Println("Equilateral")
	} else if len(numMap) == 2 {
		fmt.Println("Isosceles")
	} else if len(numMap) == 3 {
		fmt.Println("Scalene")
	}
}
