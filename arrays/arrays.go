package array

import "fmt"

func ArrayDemo() {

	var Products [3]int
	fmt.Println(Products)

	grade1 := 10
	grade2 := 20
	grade3 := 30
	grades := [3]int{grade1, grade2, grade3}
	fmt.Println(&grades[0])
	fmt.Println(&grades[1])
	fmt.Println(&grades[2])
	fmt.Println(grades)
	fmt.Println(grades[0])

	var x int8
	var y int8
	gradeall := [2]int8{x, y}
	gradeall[0] = 10
	fmt.Println(gradeall[0])

	var grad [3]int

	grad[0] = 10
	grad[1] = 20
	grad[2] = 30

	fmt.Println(grad)

	grad1 := grad
	fmt.Println(grad1)
	grad[1] = 40
	fmt.Println(grad)
	fmt.Println(grad1)
}
