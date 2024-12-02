package main

import "fmt"

func main() {
	var n, sum int
	fmt.Scan(&n)
	for i := 0; i < n-1; i++ {
		sum += i + 1
	}
	fmt.Println(sum)
	fmt.Println(2)
}
