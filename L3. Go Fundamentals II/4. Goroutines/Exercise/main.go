package main

import (
	"fmt"
	"time"
)

func iterateAndPrint(toPrint []string) {
	for _, str := range toPrint {
		fmt.Println(str)
	}
}

func main() {

	colorNames := []string{"red", "orange", "yellow", "green", "blue", "indigo", "violet"}
	colorCodes := []string{"#FF0000", "#FF7F00", "#FFFF00", "#00FF00", "#0000FF", "#4B0082", "#9400D3"}

	start := time.Now()

	go iterateAndPrint(colorNames)
	go iterateAndPrint(colorCodes)

	time.Sleep(1 * time.Second)

	elapsed := time.Since(start)

	fmt.Println("Execution time: ", elapsed)
}
