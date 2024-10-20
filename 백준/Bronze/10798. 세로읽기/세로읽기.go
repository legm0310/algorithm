package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	row := 5
	length := 0
	var result string
	reader := bufio.NewReader(os.Stdin)
	matrix := make([]string, row)
	for i := 0; i < row; i++ {
		line, _ := reader.ReadString('\n')
		values := strings.Fields(strings.TrimSpace(line))
		if length < len(strings.Join(values, "")) {
			length = len(strings.Join(values, ""))
		}
		matrix[i] = strings.Join(values, "")
	}
	for j := 0; j < length; j++ {
		for i := 0; i < row; i++ {
			if j < len(matrix[i]) {
				result += string(matrix[i][j])
			}
		}
	}
	fmt.Println(result)
}
