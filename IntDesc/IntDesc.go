package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func solution(n int64) int64 {
	str := strconv.Itoa(int(n))
	arr := strings.Split(str, "")
	sort.Sort(sort.Reverse(sort.StringSlice(arr)))
	answer, _ := strconv.Atoi(strings.Join(arr, ""))
	return int64(answer)
}

func main() {
	fmt.Println(solution(118372)) //873211
}
