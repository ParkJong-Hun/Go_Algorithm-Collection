package main

import "fmt"

func solution(arr1 [][]int, arr2 [][]int) [][]int {
    for x := range(arr1) {
        for y := range(arr1[0]) {
            arr1[x][y] += arr2[x][y]
        }
    }
    return arr1
}

func main() {
	fmt.Println(solution([[1,2],[2,3]], [[3,4],[5,6]]))
}
