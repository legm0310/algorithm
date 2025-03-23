package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
)

var (
	n, m       int
	chickenLoc [][2]int
	homeLoc    [][2]int
	cities     [50][50]int
	ans        = 1000000
	s          = bufio.NewScanner(os.Stdin)
	w          = bufio.NewWriter(os.Stdout)
)

func main() {
	defer w.Flush()
	s.Split(bufio.ScanWords)

	n, m = scanInt(), scanInt()

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			cities[i][j] = scanInt()
		}
	}

	for r := 0; r < n; r++ {
		for c := 0; c < n; c++ {
			if cities[r][c] == 1 {
				homeLoc = append(homeLoc, [2]int{r, c})
			}
			if cities[r][c] == 2 {
				chickenLoc = append(chickenLoc, [2]int{r, c})
			}
		}
	}
	arr := make([][2]int, 0)
	recur(0, &arr)
	w.WriteString(strconv.Itoa(ans) + "\n")
}

func recur(start int, places *[][2]int) {
	if len(*places) == m {
		distSum := 0
		for _, home := range homeLoc {
			dist := -1
			for _, chicken := range *places {
				if calc := getDistance(home, chicken); dist == -1 || dist > calc {
					dist = calc
				}
			}
			distSum += dist
		}
		ans = min(ans, distSum)
		return
	}

	for i := start; i < len(chickenLoc); i++ {
		*places = append(*places, chickenLoc[i])
		recur(i+1, places)
		*places = (*places)[:len(*places)-1]
	}
}

func getDistance(loc1, loc2 [2]int) int {
	xIncr := (loc1[0] + 1) - (loc2[0] + 1)
	yIncr := (loc1[1] + 1) - (loc2[1] + 1)
	return int(math.Abs(float64(xIncr)) + math.Abs(float64(yIncr)))
}

func scanInt() int {
	s.Scan()
	x, _ := strconv.Atoi(s.Text())
	return x
}
