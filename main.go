package main

import "fmt"

func PrintTitle() {
	fmt.Println("Deep Thread")
}

func Fact(n int) int {
	if n <= 0 {
		return 1
	}

	return n * Fact(n-1)
}

func FactLoop(n int) int {
	result := 1

	// Test 1
	// for n > 0 {
	// 	result *= n
	// 	n--
	// }
	// Test 2
	// var i = 2
	// for ; i <= n; i++ {
	// 	result *= i
	// }
	for i := 2; i <= n; i++ {
		result *= i
	}

	return result
}

func main() {
	PrintTitle()

	fmt.Println("Fack(5): ", Fact(5))
	fmt.Println("Fack(5): ", FactLoop(5))
}
