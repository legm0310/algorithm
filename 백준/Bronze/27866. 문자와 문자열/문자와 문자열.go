package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	r *bufio.Reader
)

func init() {
	r = bufio.NewReader(os.Stdin)
}

func main() {
	var num int
	var word string
	fmt.Fscanln(r, &word)
	fmt.Fscanln(r, &num)
	fmt.Println(string(word[num-1]))
}