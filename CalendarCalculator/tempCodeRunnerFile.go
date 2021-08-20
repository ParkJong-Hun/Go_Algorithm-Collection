package main

import "fmt"

func solution(a int, b int) string {
	day := 0
	answer := []string{"FRI", "SAT", "SUN", "MON", "TUE", "WED", "THU"}
	monthToDay := map[int]int{
		1:  0,
		2:  31,
		3:  29,
		4:  31,
		5:  30,
		6:  31,
		7:  30,
		8:  31,
		9:  31,
		10: 30,
		11: 31,
		12: 30,
	}
	for i := 1; i <= a; i++ {
		day += monthToDay[i]
	}
	day += b - 1
	return answer[day%7]
}

func main() {
	fmt.Println(solution(5, 24)) //TUE
}
