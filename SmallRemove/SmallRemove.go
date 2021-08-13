package main

import "fmt"

func solution(arr []int) []int {
	var answer []int
	var min = 1000
	for i := 0; i < len(arr); i++ {
		if min > arr[i] {
			min = arr[i]
		}
	}
	for i := 0; i < len(arr); i++ {
		if min != arr[i] {
			answer = append(answer, arr[i])
		}
	}
	if len(answer) == 0 {
		return []int{-1}
	} else {
		return answer
	}
}

func main() {
	fmt.Println(solution([]int{4, 3, 2, 1})) //[4, 3, 2]
}
