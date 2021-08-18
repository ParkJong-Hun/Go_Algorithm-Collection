package main

import "fmt"

func solution(a int, b int) int64 {
	if a > b {
		a, b = b, a
	}
	sum := 0
	for i := a; i <= b; i++ {
		sum += i
	}
	return int64(sum)
}

func main() {
	fmt.Println(solution(3, 5)) //12
}
