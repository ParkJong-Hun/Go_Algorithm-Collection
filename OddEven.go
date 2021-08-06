package main

import "fmt"

func solution(num int) string {
	if num%2 == 0 {
		return "Even"
	} else {
		return "Odd"
	}
}

func main() {
	fmt.Println(solution(3)) //"Odd"
	fmt.Println(solution(4)) //"Even"
}
