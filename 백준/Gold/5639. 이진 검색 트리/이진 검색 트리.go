package main

import (
	"bufio"
	"io"
	"os"
	"strconv"
)

var (
	root *Node
	s    = bufio.NewScanner(os.Stdin)
	w    = bufio.NewWriter(os.Stdout)
)

type Node struct {
	Value       int
	Left, Right *Node
}

func main() {
	defer w.Flush()
	s.Split(bufio.ScanWords)
	for {
		num, err := scanInt()
		if err == io.EOF {
			break
		}
		root = insert(root, num)
	}
	postOrder(root)
	w.WriteByte('\n')
}

func postOrder(n *Node) {
	if n == nil {
		return
	}
	postOrder(n.Left)
	postOrder(n.Right)
	w.WriteString(strconv.Itoa(n.Value) + "\n")

}

func insert(n *Node, v int) *Node {
	if n == nil {
		return &Node{Value: v}
	}
	if v < n.Value {
		n.Left = insert(n.Left, v)
	} else {
		n.Right = insert(n.Right, v)
	}
	return n
}

func scanInt() (int, error) {
	if !s.Scan() {
		return 0, io.EOF
	}
	x, _ := strconv.Atoi(s.Text())
	return x, nil
}
