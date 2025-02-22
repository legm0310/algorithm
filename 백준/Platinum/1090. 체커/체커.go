package main

import (
	"bufio"
	"math"
	"os"
	"slices"
	"strconv"
)

var (
	coords [][]int
	s      *bufio.Scanner
	w      *bufio.Writer
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

func main() {
	defer w.Flush()
	N, _ := scanInt()
	xArr := make([]int, N)
	yArr := make([]int, N)
	result := make([]int, N)
	coords = make([][]int, N)
	for i := 0; i < N; i++ {
		result[i] = -1
		coords[i] = make([]int, 2)
		coords[i][0], _ = scanInt()
		coords[i][1], _ = scanInt()
		xArr[i] = coords[i][0]
		yArr[i] = coords[i][1]
	}

	for _, y := range yArr {
		for _, x := range xArr {
			dist := make([]int, N)
			for i, coord := range coords {
				dist[i] = int(math.Abs(float64(coord[0]-x)) + math.Abs(float64(coord[1]-y)))
			}

			slices.Sort(dist)
			tmp := 0
			for i, d := range dist {
				if i >= N {
					break
				}
				tmp += d
				if result[i] == -1 {
					result[i] = tmp
				} else {
					result[i] = min(result[i], tmp)
				}
			}
		}
	}

	for i := 0; i < N; i++ {
		w.WriteString(strconv.Itoa(result[i]))
		w.WriteString(" ")
	}
	w.WriteByte('\n')
}
