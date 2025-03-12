package main

import (
	"bufio"
	"os"
	"strconv"
)

var (
	a, b int
	s    *bufio.Scanner
	w    *bufio.Writer
)

func init() {
	s = bufio.NewScanner(os.Stdin)
	w = bufio.NewWriter(os.Stdout)
	s.Split(bufio.ScanWords)
}

func scanInt() (int, error) {
	s.Scan()
	return strconv.Atoi(s.Text())
}

func dfs(start int) int {
	queue := []int{start}
	dist := make(map[int]int)
	dist[start] = 1

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		if cur == b {
			return dist[cur]
		}

		next := []int{cur * 2, cur*10 + 1}
		for _, n := range next {
			if n <= b {
				if _, exists := dist[n]; !exists {
					dist[n] = dist[cur] + 1
					queue = append(queue, n)
				}
			}
		}
	}

	return -1
}

func main() {
	defer w.Flush()
	a, _ = scanInt()
	b, _ = scanInt()
	cnt := 1
	for a < b {
		if b%10 == 1 {
			b /= 10
		} else if b%2 == 0 {
			b /= 2
		} else {
			w.WriteString(strconv.Itoa(-1) + "\n")
			return
		}
		cnt++
	}

	if a == b {
		w.WriteString(strconv.Itoa(cnt) + "\n")
	} else {
		w.WriteString(strconv.Itoa(-1) + "\n")
	}
}
