package main

import (
	"fmt"
	"strings"
)

func solution(s string) string {
	var answer string
	index := 0
	for _, word := range s {
		char := string(word)
		if char == " " {
			index = 0
			answer += char
			continue
		}
		if index%2 == 0 {
			answer += strings.ToUpper(char)
		} else {
			answer += strings.ToLower(char)
		}
		index++
	}
	return answer
}

func main() {
	fmt.Println(solution("try hello world")) //"TrY HeLlO WoRlD"
}
