package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var maxNum, minNum int
	minPrime := 0
	sum := 0
	fmt.Fscanf(reader, "%d\n%d\n", &minNum, &maxNum)
	for i := minNum; i <= maxNum; i++ {
		target := int(math.Sqrt(float64(i)))
		isPrime := true
		if i <= 1 {
			continue
		}
		for j := 2; j <= target; j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}

		if isPrime {
			if minPrime == 0 {
				minPrime = i
			} else if i < minPrime {
				minPrime = i
			}
			sum += i
		}
	}
	if sum == 0 {
		fmt.Println(-1)
	} else {
		fmt.Printf("%d\n%d\n", sum, minPrime)
	}
}
