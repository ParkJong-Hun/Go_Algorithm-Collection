package main

import "fmt"

func gcd(num1 int, num2 int) int {
	if num2 == 0 {
		return num1
	} else {
		return gcd(num2, num1%num2)
	}
}

func solution(n int, m int) []int {
	var answer []int
	answer = append(answer, gcd(n, m))
	answer = append(answer, n*m/gcd(n, m))
	return answer
}

func main() {
	fmt.Println(solution(3, 12)) //3, 12
	fmt.Println(solution(2, 5))  //1, 10
}
