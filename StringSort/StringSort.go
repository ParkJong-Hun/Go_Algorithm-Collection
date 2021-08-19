package main

import (
	"fmt"
	"sort"
)

func solution(strings []string, n int) []string {
	sort.Slice(strings, func(i, j int) bool {
		if strings[i][n] == strings[j][n] {
			return strings[i] < strings[j]
		}
		return strings[i][n] < strings[j][n]
	})
	return strings
}

func main() {
	fmt.Println(solution([]string{"sun", "bed", "car"}, 1))
}
