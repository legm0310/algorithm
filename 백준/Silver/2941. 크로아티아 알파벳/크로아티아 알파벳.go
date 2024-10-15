package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var word string
	var arr []byte
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &word)
	word = strings.Replace(word, "c=", "1", -1)
	word = strings.Replace(word, "c-", "1", -1)
	word = strings.Replace(word, "dz=", "1", -1)
	word = strings.Replace(word, "d-", "1", -1)
	word = strings.Replace(word, "lj", "1", -1)
	word = strings.Replace(word, "nj", "1", -1)
	word = strings.Replace(word, "s=", "1", -1)
	word = strings.Replace(word, "z=", "1", -1)
	for i := 0; i< len(word); i++ {
		arr = append(arr, (word[i]))
	}
	fmt.Println(len(string(arr)))
}