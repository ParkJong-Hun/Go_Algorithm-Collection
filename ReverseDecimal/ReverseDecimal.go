package main

import "fmt"

//이런식으로 return 변수를 만들 수 있음
func solution(n int64) (result []int) {
	for n > 0 {
		result = append(result, int(n%10))
		n = int64(n / 10)
	}
	return
}

func main() {
	fmt.Println(solution(12345)) //[5, 4, 3, 2, 1]
}
