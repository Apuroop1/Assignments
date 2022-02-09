package array

import "fmt"

func ArrayDemo() {

	var Products [3]int
	fmt.Println(Products)

	grade1 := 10
	grade2 := 20
	grade3 := 30
	grades := [3]int{grade1, grade2, grade3}
	fmt.Println(grades)
	fmt.Println(grades[0])
}
