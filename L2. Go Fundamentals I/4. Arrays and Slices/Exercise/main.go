package main

import "fmt"

func main() {
	// 0 1 1 2 3 5 8 13

	fibSlice := []int{0, 1, 1, 2, 3, 5, 8, 13}

	fmt.Println(fibSlice)
	fmt.Println(len(fibSlice))
	fmt.Println(fibSlice[0:4])

	fibSlice = append(fibSlice, 21)

	fmt.Println(fibSlice)
	fmt.Println(len(fibSlice))
}
