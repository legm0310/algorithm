package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	ints := make([]int, 3)
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanf(reader, "%d %d %d\n", &ints[0], &ints[1], &ints[2])
	sort.Ints(ints)
	if ints[0]+ints[1] > ints[2] {
		fmt.Println(ints[0] + ints[1] + ints[2])
	} else {
		fmt.Println(2*(ints[0]+ints[1]) - 1)
	}
}
