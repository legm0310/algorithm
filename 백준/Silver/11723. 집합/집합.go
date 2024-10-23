package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Set struct {
	bits uint32
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	set := &Set{}

	var output strings.Builder

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) == 0 {
			break
		}
		method := parts[0]
		var num int
		if len(parts) == 2 {
			num, _ = strconv.Atoi(parts[1])
		}
		switch method {
		case "add":
			set.bits |= (1 << num)
		case "remove":
			set.bits &= ^(1 << num)
		case "check":
			if set.bits&(1<<num) != 0 {
				output.WriteString("1\n")
			} else {
				output.WriteString("0\n")
			}
		case "toggle":
			set.bits ^= (1 << num)
		case "all":
			set.bits = (1 << 21) - 1
		case "empty":
			set.bits = 0
		}
	}
	fmt.Print(output.String())
}
