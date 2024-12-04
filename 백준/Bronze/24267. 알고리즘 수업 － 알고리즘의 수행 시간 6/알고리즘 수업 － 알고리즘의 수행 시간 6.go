package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(n * (n - 1) * (n - 2) / 6)
	fmt.Println(3)
}