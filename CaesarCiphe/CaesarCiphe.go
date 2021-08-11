package main

import "fmt"

func solution(s string, n int) string {
	//rune: UTF-8을 표현하는 타입
	answer := make([]rune, len(s))
	offset := rune(n)
	for i, v := range s {
		if 'A' <= v && v <= 'Z' {
			answer[i] = ((v + offset - 'A') % 26) + 'A'
		} else if 'a' <= v && v <= 'z' {
			answer[i] = ((v + offset - 'a') % 26) + 'a'
		} else {
			answer[i] = v
		}
	}
	return string(answer)
}

func main() {
	fmt.Println(solution("AB", 1)) //BC
}
