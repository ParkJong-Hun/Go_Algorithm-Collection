package main

import "fmt"

func solution(s string) bool {
	if len(s) != 4 && len(s) != 6 {
		return false
	}
	for _, char := range s {
		if char > '9' || char < '0' {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(solution("a234")) //false
	fmt.Println(solution("1234")) //true
}
