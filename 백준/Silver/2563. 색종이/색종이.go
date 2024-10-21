package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	length := 10
	var paperCnt int
	fmt.Fscanln(reader, &paperCnt)
	unique := make(map[string]struct{})
	for k := 0; k < paperCnt; k++ {
		var colStart, rowStart int
		fmt.Fscanln(reader, &colStart, &rowStart)
		for i := 0; i < length; i++ {
			for j := 0; j < length; j++ {
				key := strconv.Itoa(j+colStart) + "," + strconv.Itoa(i+rowStart)
				unique[key] = struct{}{}

			}
		}
	}
	fmt.Printf(strconv.Itoa(len(unique)))
}
