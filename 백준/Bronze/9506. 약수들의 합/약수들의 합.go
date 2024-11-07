package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		var n, tmp int
		fmt.Fscanf(reader, "%d\n", &n)
		if n == -1 {
			break
		}
		s := make([]int, 0, n)
		for i := 1; i < n; i++ {
			if n%i == 0 {
				s = append(s, i)
				tmp += i
			}
		}
		if tmp == n {
			fmt.Printf("%d = ", n)
			for i := 0; i < len(s); i++ {
				if i == 0 {
					fmt.Printf("%d ", s[i])
				} else if i == len(s)-1 {
					fmt.Printf("+ %d\n", s[i])
				} else {
					fmt.Printf("+ %d ", s[i])
				}
			}
		} else {
			fmt.Println(n, "is NOT perfect.")
		}

	}
}