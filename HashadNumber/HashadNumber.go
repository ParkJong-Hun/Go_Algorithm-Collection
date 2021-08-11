package main

import (
	"fmt"
)

func solution(x int) bool {
	n := x
	sum := 0
	for n != 0 {
		sum += n % 10
		n /= 10
	}
	return x%sum == 0
}

func main() {
	fmt.Println(solution(10)) //true
	fmt.Println(solution(12)) //true
	fmt.Println(solution(11)) //true
	fmt.Println(solution(13)) //true
}
