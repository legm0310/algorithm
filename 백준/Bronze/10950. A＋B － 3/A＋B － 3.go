package main

import (
	"bufio"
	"fmt"
	"os"
)	

func main() {
	reader := bufio.NewReader(os.Stdin)
	var length int
    fmt.Fscanln(reader, &length)
    for i:=0 ; i<length; i++ {
        var num1, num2 int
        fmt.Fscanln(reader, &num1, &num2)
        fmt.Println(num1+num2)
    }
}