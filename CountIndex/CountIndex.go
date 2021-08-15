package main

import (
	"fmt"
	"strconv"
)

func solution(n int) int {
	answer := 0
	s := strconv.Itoa(n)
	for _, v := range s {
		answer += (int(v) - '0')
	}
	return answer
}

func main() {
	fmt.Println(solution(123)) //6
	fmt.Println(solution(987)) //24
}
