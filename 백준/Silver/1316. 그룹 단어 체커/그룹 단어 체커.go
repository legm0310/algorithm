package main

import (
	"bufio"
	"fmt"
	"os"
)

func isGroup(word string) bool {
	var chars = make(map[rune]bool)
	var prev rune

	for _, char := range word {
		if char != prev {
			if chars[char] {
				return false
			}
		}
		chars[char] = true
		prev = char
	}
	return true
}

func main() {
	var n int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &n)
	
	var result int
	var word string
	for i := 0; i<n; i++ {
		fmt.Fscanln(reader, &word)
		if isGroup(word) {
			result++
		}
	}
	fmt.Println(result)
}