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
	var num int
	var base int
	fmt.Fscanf(reader, "%d %d", &num, &base)
	result := strconv.FormatInt(int64(num), base)
	fmt.Println(strings.ToUpper(result))
}