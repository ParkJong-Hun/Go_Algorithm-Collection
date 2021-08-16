package main

import "fmt"

func solution(seoul []string) string {
	for index, name := range seoul {
		if name == "Kim" {
			return fmt.Sprintf("김서방은 %d에 있다", index)
		}
	}
	return ""
}

func main() {
	fmt.Println(solution([]string{"Jane", "Kim"})) //1
}
