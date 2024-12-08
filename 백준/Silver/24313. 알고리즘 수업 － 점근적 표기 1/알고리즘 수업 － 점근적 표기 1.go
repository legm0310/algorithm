package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var a1, a0, c, n0 int
	r := bufio.NewReader(os.Stdin)
	fmt.Fscanln(r, &a1, &a0)
	fmt.Fscanln(r, &c)
	fmt.Fscanln(r, &n0)
	if a1 < c {
		if n0*a1+a0 <= c*n0 {
			fmt.Println(1)
		} else {
			fmt.Println(0)
		}
	} else if a1 > c {
		fmt.Println(0)
	} else {
		if a0 <= 0 {
			if n0*a1+a0 <= c*n0 {
				fmt.Println(1)
			} else {
				fmt.Println(0)
			}
		} else {
			fmt.Println(0)
		}
	}

}
