package main

import (
	"bufio"
	"os"
	"slices"
	"strings"
)

var (
	s *bufio.Scanner
	w *bufio.Writer
)

func init() {
	s = bufio.NewScanner(os.Stdin)
	w = bufio.NewWriter(os.Stdout)
	s.Split(bufio.ScanWords)
}

func main() {
	defer w.Flush()
	s.Scan()
	N := s.Text()
	s.Scan()
	N2 := s.Text()
	nums := make([]string, 3)
	nums2 := make([]string, 3)

	for i := 0; i < 3; i++ {
		nums[i] = string(N[i])
		nums2[i] = string(N2[i])
	}
	slices.Reverse(nums)
	slices.Reverse(nums2)
	var result string
	if strings.Join(nums, "") > strings.Join(nums2, "") {
		result = strings.Join(nums, "")
	} else {
		result = strings.Join(nums2, "")
	}
	w.WriteString(result)
	w.WriteByte('\n')
}