package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	var n, m int
	fmt.Fscanln(r, &n, &m)

	result := make([]int, 0)
	for i := 0; i < 2; i++ {
		line, _ := r.ReadString('\n')
		line = strings.TrimSpace(line)
		parts := strings.Fields(line)
		for _, v := range parts {
			num, _ := strconv.Atoi(v)
			result = append(result, num)
		}
	}

	sort.Ints(result)
	fmt.Println(strings.Trim(fmt.Sprint(result), "[]"))
}
