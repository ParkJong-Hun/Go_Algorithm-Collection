package main

import (
	"fmt"
	"math"
)

func solution(n int64) int64 {
	sqrt := int64(math.Sqrt(float64(n)))
	if sqrt*sqrt == n {
		sqrt++
		return sqrt * sqrt
	}
	return -1
}

func main() {
	fmt.Println(solution(121)) //144
	fmt.Println(solution(3))   //-1
}
