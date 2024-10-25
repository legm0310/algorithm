package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	parts := strings.Fields(scanner.Text())
	base, _ := strconv.Atoi(parts[1])
	result, _ := strconv.ParseInt(parts[0], base, 64)
	fmt.Println(result)
}