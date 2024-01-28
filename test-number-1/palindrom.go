package main

import "fmt"

func isPalindrome(str string) bool {
	length := len(str)
	i, j := 0, length-1

	for i < j {
		if str[i] != str[j] {
			return false
		}
		i++
		j--
	}

	return true
}

func main() {
	input1 := "radar"
	input2 := "hello"
	fmt.Printf("Input: %s\nOutput: %t\n\n", input1, isPalindrome(input1))
	fmt.Printf("Input: %s\nOutput: %t", input2, isPalindrome(input2))
}
