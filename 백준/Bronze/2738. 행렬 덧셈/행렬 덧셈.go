package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var rows int
	var cols int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &rows, &cols)
	matrix := make([][]int, rows)
	for i := 0; i < rows; i++ {
		matrix[i] = make([]int, cols)
	}
	for i := 0; ; i++ {
		currentRow := i % rows
		var line string
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
		}
		if line == "" {
			break
		}
		values := strings.Fields(line)
		for j := 0; j < len(values); j++ {
			num, _ := strconv.Atoi(values[j])
			matrix[currentRow][j] += num
		}
	}
	for i := 0; i < rows; i++ {
		strArr := make([]string, len(matrix[i]))
		for j, v := range matrix[i] {
			strArr[j] = strconv.Itoa(v)
		}
		fmt.Println(strings.Join(strArr, " "))
	}
}
