package main

import (
	"bufio"
	"os"
	"strconv"
)

const INF = 100000000

var (
	n, m int
	dist [][]int
	s    = bufio.NewScanner(os.Stdin)
	w    = bufio.NewWriter(os.Stdout)
)

func main() {
	defer w.Flush()
	s.Split(bufio.ScanWords)
	n, m = scanInt(), scanInt()
	dist = make([][]int, n+1)
	for i := 1; i < n+1; i++ {
		dist[i] = make([]int, n+1)
		for j := 1; j <= n; j++ {
			if i == j {
				dist[i][j] = 0
			} else {
				dist[i][j] = INF
			}
		}
	}
	for i := 0; i < m; i++ {
		a, b, c := scanInt(), scanInt(), scanInt()
		if c < dist[a][b] {
			dist[a][b] = c
		}
	}

	for k := 1; k <= n; k++ {
		for i := 1; i <= n; i++ {
			for j := 1; j <= n; j++ {
				if dist[i][j] > dist[i][k]+dist[k][j] {
					dist[i][j] = dist[i][k] + dist[k][j]
				}
			}
		}
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if dist[i][j] == INF {
				w.WriteString(strconv.Itoa(0) + " ")
			} else {
				w.WriteString(strconv.Itoa(dist[i][j]) + " ")
			}
		}
		w.WriteByte('\n')
	}

}

func scanInt() int {
	s.Scan()
	x, _ := strconv.Atoi(s.Text())
	return x
}
