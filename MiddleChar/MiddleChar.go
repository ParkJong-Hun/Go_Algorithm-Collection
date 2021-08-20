package main

import "fmt"

func solution(s string) string {
	return s[int(float64(len(s))/2-0.5):int(float64(len(s))/2+1)]
}

func main() {
	fmt.Println(solution("abcde")) //c
	fmt.Println(solution("qwer"))  //we
}
