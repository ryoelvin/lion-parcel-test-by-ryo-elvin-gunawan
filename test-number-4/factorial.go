package main

import "fmt"

func factorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}

	return n * factorial(n-1)
}

func main() {
	input := 5
	fmt.Printf("Input: %d\nOutput: %d", input, factorial(input))
}
