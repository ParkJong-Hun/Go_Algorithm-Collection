package main

import "fmt"

func solution(n int) string {
	var answer string = ""
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			answer += "수"
		} else {
			answer += "박"
		}
	}
	return answer
}

func main() {
	fmt.Println(solution(3)) //수박수
	fmt.Println(solution(4)) //수박수박
}
