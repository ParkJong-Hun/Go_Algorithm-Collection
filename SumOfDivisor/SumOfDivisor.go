package main

import "fmt"

func solution(n int) (sum int) {
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			sum += i
		}
	}
	return
}

func main() {
	fmt.Println(solution(12)) //28
	fmt.Println(solution(5))  //6
}
