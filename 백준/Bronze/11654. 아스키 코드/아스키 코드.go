package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	s *bufio.Scanner
)

func init() {
	s = bufio.NewScanner(os.Stdin)
}

func main() {
	s.Scan()
	r := []rune(s.Text())
	fmt.Printf("%d\n", r[0])
}