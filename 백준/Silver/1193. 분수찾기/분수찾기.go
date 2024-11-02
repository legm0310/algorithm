package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscanf(reader, "%d\n", &n)
	cnt := 1
	for i := 1; ; i++ {
		exit := false
		tmp := i
		for j := 0; j < i; j++ {
			if i%2 == 0 {
				if cnt == n {
					fmt.Println(strconv.Itoa(j+1) + "/" + strconv.Itoa(tmp))
					exit = true
					break
				}
				tmp--
			}
			if i%2 != 0 {
				if cnt == n {
					fmt.Println(strconv.Itoa(tmp) + "/" + strconv.Itoa(j+1))
					exit = true
					break
				}
				tmp--
			}
			cnt++
		}
		if exit {
			break
		}

	}

}