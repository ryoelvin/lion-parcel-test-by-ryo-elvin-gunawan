package main

import "fmt"

func printRightTriangle(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		if i != n {
			fmt.Println()
		}
	}
}

func main() {
	input := 5
	fmt.Printf("Input: %d\n\nOutput:\n", input)
	printRightTriangle(input)
}
