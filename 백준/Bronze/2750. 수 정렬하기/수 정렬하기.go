package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	var cnt int
	var nums []int
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	fmt.Fscanln(r, &cnt)
	for i := 0; i < cnt; i++ {
		var input int
		isContain := false
		fmt.Fscanln(r, &input)
		for _, num := range nums {
			if num == input {
				isContain = true
				break
			}
		}
		if isContain == true {
			continue
		}
		nums = append(nums, input)
	}
	sort.Ints(nums)
	for _, num := range nums {
		w.WriteString(strconv.Itoa(num) + "\n")
	}
}
