package main

import (
	"fmt"
)

// Uncomment the line below after implementing your reduce() function

// TODO: create a reduce() function here

func main() {
	// Uncomment the line below after implementing your reduce() function
	fmt.Println(reduce([]int{0, 1, 1, 2, 3, 5, 8, 13, 21}))
}

func reduce(numbers []int) int {
	sum := 0
	for _, n := range numbers {
		sum += n
	}
	return sum
}
