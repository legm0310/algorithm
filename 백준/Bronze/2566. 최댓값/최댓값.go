package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var result int = 0
	maxRow := 1
	maxCol := 1
	reader := bufio.NewReader(os.Stdin)
	for i := 0; i < 9; i++ {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		values := strings.Fields(line)
		for j := 0; j < 9; j++ {
			num, _ := strconv.Atoi(values[j])
			if num > result {
				result = num
				maxRow = i + 1
				maxCol = j + 1
			}
		}
	}
	fmt.Println(result)
	fmt.Println(maxRow, maxCol)
}
