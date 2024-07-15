package main

import (
    "fmt"
    s "strings"
    u "unicode"
)

func main() {
    var s1 string
    fmt.Scan(&s1)
    for _, ch := range s1 {
        if u.IsUpper(ch) {
            fmt.Printf("%s",s.ToLower(string(ch)))
            
        } else {
            fmt.Printf("%s", s.ToUpper(string(ch)))
        }
    }
}