package main

import "fmt"

func findMax(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	max := arr[0]

	for _, num := range arr {
		if num > max {
			max = num
		}
	}

	return max
}

func main() {
	input := []int{3, 5, 1, 9, 2}
	fmt.Printf("Input: %v\nOutput: %d", input, findMax(input))
}
