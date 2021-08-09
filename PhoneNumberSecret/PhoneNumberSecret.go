package main

import "fmt"

func solution(phone_number string) string {
	var answer string
	for i := 0; i < len(phone_number); i++ {
		if i < len(phone_number)-4 {
			answer += "*"
		} else {
			answer += string(phone_number[i])
		}
	}
	return answer
}

func main() {
	fmt.Println(solution("01033334444"))
}
