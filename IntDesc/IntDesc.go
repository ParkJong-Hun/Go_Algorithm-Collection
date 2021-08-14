package main

import (
	"fmt"
	"strconv"
	"strings"
)

func solution(n int64) int64 {
	str := strconv.Itoa(int(n))
	arr := strings.Split(str, "")
	answer, _ := strconv.Atoi(strings.Join(arr, ""))
	return int64(answer)
}

func main() {
	fmt.Println(solution(118372)) //873211
}
