package main

import (
	"fmt"
	"sort"
)

func solution(arr []int, divisor int) []int {
	var returnArray []int
	for _, i := range arr {
		if i%divisor == 0 {
			returnArray = append(returnArray, i)
		}
	}
	sort.Ints(returnArray)
	if len(returnArray) == 0 {
		returnArray = append(returnArray, -1)
	}
	return returnArray
}

func main() {
	fmt.Println(solution([]int{5, 9, 7, 10}, 5)) //5, 10
}
