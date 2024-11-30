package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	for {
		ns := make([]int, 3)
		fmt.Fscanln(r, &ns[0], &ns[1], &ns[2])
		if ns[0] == 0 {
			return
		}
		sort.Ints(ns)
		if ns[2] >= ns[0]+ns[1] {
			w.WriteString("Invalid\n")
		} else if ns[0] == ns[1] && ns[1] == ns[2] {
			w.WriteString("Equilateral\n")
		} else if ns[0] == ns[1] || ns[1] == ns[2] {
			w.WriteString("Isosceles\n")
		} else if ns[0] != ns[1] && ns[1] != ns[2] && ns[0] != ns[2] {
			w.WriteString("Scalene\n")
		}
	}
}
