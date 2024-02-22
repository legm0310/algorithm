package main

import "fmt"

func main() {
    var s1 string
    fmt.Scan(&s1)
    for idx, _ := range s1 {
        fmt.Println(string(s1[idx]))
    }
}