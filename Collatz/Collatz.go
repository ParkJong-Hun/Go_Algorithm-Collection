package main

import "fmt"

func solution(num int) int {
	count := 0
	result := num
	for result != 1 {
		if count >= 500 {
			return -1
		}
		if result%2 == 0 {
			result /= 2
		} else {
			result = result*3 + 1
		}
		count += 1
	}
	return count
}

func main() {
	fmt.Println(solution(6))      //8
	fmt.Println(solution(16))     //4
	fmt.Println(solution(626331)) //-1
}
