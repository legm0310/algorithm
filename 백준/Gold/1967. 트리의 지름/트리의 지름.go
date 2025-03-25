package main

import (
	"bufio"
	"os"
	"strconv"
)

var (
	n, farthest int
	tree        [][]*Edge
	ans         = 0
	s           = bufio.NewScanner(os.Stdin)
	w           = bufio.NewWriter(os.Stdout)
)

type Edge struct {
	v, w int
}

func main() {
	defer w.Flush()
	s.Split(bufio.ScanWords)
	n = scanInt()
	tree = make([][]*Edge, n+1)
	for i := 0; i < n-1; i++ {
		parent, child, weight := scanInt(), scanInt(), scanInt()
		tree[parent] = append(tree[parent], &Edge{child, weight})
		tree[child] = append(tree[child], &Edge{parent, weight})
	}

	find(1, 0, 0)
	ans = 0
	find(farthest, 0, 0)
	w.WriteString(strconv.Itoa(ans))
	w.WriteByte('\n')
}

func find(v, parent, dist int) {
	if dist > ans {
		ans = dist
		farthest = v
	}
	for _, nxt := range tree[v] {
		if nxt.v != parent {
			find(nxt.v, v, dist+nxt.w)
		}
	}
}

func scanInt() int {
	s.Scan()
	x, _ := strconv.Atoi(s.Text())
	return x
}
