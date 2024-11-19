package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func findUnique(n []int) int {
	if n[0] == n[1] {
		return n[2]
	} else if n[1] == n[2] {
		return n[0]
	} else {
		return n[1]
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	xSlice := make([]int, 3)
	ySlice := make([]int, 3)
	for i := 0; i < 3; i++ {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		parts := strings.Fields(line)
		xSlice[i], _ = strconv.Atoi(parts[0])
		ySlice[i], _ = strconv.Atoi(parts[1])
	}
	fmt.Println(findUnique(xSlice), findUnique(ySlice))
}
