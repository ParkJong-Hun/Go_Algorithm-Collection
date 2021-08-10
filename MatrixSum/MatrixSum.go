package main

import "fmt"

func solution(arr1 [][]int, arr2 [][]int) [][]int {
	var answer [][]int
	for i := range arr1 {
		var sum []int
		for j := range arr1[0] {
			sum = append(sum, arr1[i][j]+arr2[i][j])
		}
		answer = append(answer, sum)
	}
	return answer
}

func main() {
	fmt.Printf("%d", solution([][]int{{1, 2}, {2, 3}}, [][]int{{3, 4}, {5, 6}}))
}
