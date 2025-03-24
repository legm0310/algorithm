package main

import (
	"bufio"
	"os"
	"strconv"
)

var (
	n      int
	parent []int
	//tree   = map[int][]int{}
	tree [][]int
	s    = bufio.NewScanner(os.Stdin)
	w    = bufio.NewWriter(os.Stdout)
)

func main() {
	defer w.Flush()
	s.Split(bufio.ScanWords)
	n = scanInt()
	parent = make([]int, n+1)
	tree = make([][]int, n+1)
	for i := 0; i < n-1; i++ {
		a, b := scanInt(), scanInt()
		tree[a] = append(tree[a], b)
		tree[b] = append(tree[b], a)
	}
	recur(1, 0)
	for _, i := range parent {
		if i != 0 {
			w.WriteString(strconv.Itoa(i) + "\n")
		}
	}
}

func recur(node, prev int) {
	parent[node] = prev

	for _, nxt := range tree[node] {
		if nxt == prev {
			continue
		} else {
			recur(nxt, node)
		}
	}
}

func scanInt() int {
	s.Scan()
	x, _ := strconv.Atoi(s.Text())
	return x
}
