package main

import (
	"bufio"
	"os"
	"strconv"
)

var (
	n             int
	pre, in, post []byte
	tree          = map[byte][2]byte{}
	s             = bufio.NewScanner(os.Stdin)
	w             = bufio.NewWriter(os.Stdout)
)

func main() {
	defer w.Flush()
	s.Split(bufio.ScanWords)
	n = scanInt()
	for i := 0; i < n; i++ {
		node, left, right := next(), next(), next()
		tree[node] = [2]byte{left, right}
	}
	recur('A')
	w.WriteString(string(pre) + "\n")
	w.WriteString(string(in) + "\n")
	w.WriteString(string(post) + "\n")
}

func recur(node byte) {
	if node == '.' {
		return
	}
	pre = append(pre, node)
	recur(tree[node][0])
	in = append(in, node)
	recur(tree[node][1])
	post = append(post, node)
}

func next() byte {
	s.Scan()
	x := s.Bytes()
	return x[0]
}

func scanInt() int {
	s.Scan()
	x, _ := strconv.Atoi(s.Text())
	return x
}
