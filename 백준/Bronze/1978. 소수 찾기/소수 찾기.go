package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var idx, cnt int
	fmt.Fscanf(reader, "%d\n", &idx)
	for i := 0; i < idx; i++ {
		flag := true
		var n int
		fmt.Fscanf(reader, "%d ", &n)
		sqrt := int(math.Sqrt(float64(n)))
		if n <= 1 {
			continue
		}
		for j := 2; j <= sqrt; j++ {
			if n%j == 0 {
				flag = false
				break
			}
		}
		if flag {
			cnt++
		}
	}
	fmt.Println(cnt)
}
