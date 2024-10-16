package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scoreTable := map[string]float32{
		"A+": 4.5,
		"A0": 4.0,
		"B+": 3.5,
		"B0": 3.0,
		"C+": 2.5,
		"C0": 2.0,
		"D+": 1.5,
		"D0": 1.0,
		"F": 0.0,
	}
	// 전공 평점
	var GPA float32
	// 학점 총합
	var sumCredit float32
	var className, score string
	var credit float32
	reader := bufio.NewReader(os.Stdin)
	
	for {
		_, err := fmt.Fscanln(reader, &className, &credit, &score)
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			
		}
		if score == "P" { continue }
		sumCredit += credit
		GPA = GPA + credit * (scoreTable[score])
	}
	GPA = GPA / sumCredit
	fmt.Println(GPA)
}