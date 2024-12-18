package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var n, k int
	r := bufio.NewReader(os.Stdin)
	fmt.Fscanln(r, &n, &k)
	var nums = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &nums[i])
	}
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	fmt.Println(nums[k-1])
}
