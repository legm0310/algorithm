package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var num int64
	var base int
	fmt.Fscanf(reader, "%d %d", &num, &base)
	result := strconv.FormatInt(num, base)
	fmt.Println(strings.ToUpper(result))
}