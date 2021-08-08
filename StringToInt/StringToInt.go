package main

import (
	"fmt"
	"strconv"
)

func solution(s string) int {
	if answer, error := strconv.Atoi(s); error == nil {
		return answer
	}
	return 0
}

func main() {
	fmt.Println(solution("+1234")) //1234
	fmt.Println(solution("1234"))  //1234
	fmt.Println(solution("-1234")) //-1234
}
