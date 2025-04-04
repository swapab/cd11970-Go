package main

import (
	"fmt"
	"strings"
)

func main() {
	courses := map[uint16]string{
		1: "Calculus",
		2: "Biology",
		3: "Chemistry",
		4: "Computer Science",
		5: "Communications",
		6: "English",
		7: "Cantonese",
	}

	fmt.Println(courses)

	for _, course := range courses {
		if strings.HasPrefix(course, "C") {
			println(course)
		}
	}

	courses[4] = "Algorithms"

	courses[8] = "Spanish"

	delete(courses, 1)

	fmt.Println(courses)
}
