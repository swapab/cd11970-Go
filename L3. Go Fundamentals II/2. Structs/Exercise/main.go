package main

import "fmt"

type Student struct {
	id   uint
	name string
}

type Classroom struct {
	id          uint
	capacity    uint16
	subject     string
	studentList []Student
}

func main() {
	student1 := Student{id: 1, name: "Ramesh"}
	student2 := Student{id: 2, name: "Suresh"}
	students := []Student{student1, student2}
	class1 := Classroom{id: 1, capacity: 20, subject: "Computer Science", studentList: students}
	class2 := new(Classroom)
	class2.id = 2
	class2.capacity = 50
	class2.subject = "Mathematics"
	class2.studentList = []Student{
		{
			id:   3,
			name: "Sita",
		},
		{
			id:   4,
			name: "Gita",
		},
	}

	fmt.Println(class1)
	fmt.Println(class2)
}
