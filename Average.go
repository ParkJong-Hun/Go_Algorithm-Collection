package main

func solution(arr []int) float64 {
	sum := 0
	for _, value := range arr {
		sum += value
	}
	return float64(sum) / float64(len(arr))
}

func main() {
	numbers := []int{1, 2, 3, 4}
	solution(numbers) //2.5
}
