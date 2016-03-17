package main

import "fmt"

func main() {

	//재귀호출
	fmt.Println(factorial(5))
}

func factorial(n int) int {
	if n == 1 {
		return 1
	}

	return n * factorial(n-1)
}
