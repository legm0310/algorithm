package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n, m, result int
	r := bufio.NewReader(os.Stdin)
	fmt.Fscanln(r, &n, &m)
	line, _ := r.ReadString('\n')
	line = strings.TrimSpace(line)
	nums := make([]int, n)
	for i, numStr := range strings.Fields(line) {
		num, _ := strconv.Atoi(numStr)
		nums[i] = num
	}
	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			for k := j + 1; k < n; k++ {
				if m >= nums[i]+nums[j]+nums[k] {
					if m-result > m-(nums[i]+nums[j]+nums[k]) {
						result = nums[i] + nums[j] + nums[k]
					}
				}
			}
		}
	}
	fmt.Println(result)
}
