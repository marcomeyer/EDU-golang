package main

import "fmt"

func main() {
	
	seq := createRange(10)

	for _, n := range seq {
		fmt.Printf("%v is %v\n",n, getParityOf(n))
	}
}

func getParityOf(n int) string {
	
	if n % 2 == 0 {
		return "even"
	} else {
		return "odd"
	}
}

func createRange(n int) []int {
	
	if n < 0 {
		return []int { }
	}

	return append(createRange(n-1), n)
}