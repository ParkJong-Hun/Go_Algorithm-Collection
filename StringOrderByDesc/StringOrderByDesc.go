package main

import (
	"fmt"
	"sort"
	"strings"
)

func solution(s string) string {
	str := strings.Split(s, "")
	sort.Sort(sort.Reverse(sort.StringSlice(str)))
	return strings.Join(str, "")
}

func main() {
	fmt.Println(solution("Zbcdefg")) //"gfedcbZ"
}
