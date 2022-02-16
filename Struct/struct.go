package Struct

import "fmt"

type Movie struct {
	Name          string
	YearOfRelease int
	Actors        []string
	Grade         []int
	AvgOfGrade    int
}

func StructDemo() {

	//var m1 Movie
	m1 := new(Movie)
	m1.Name = "Apuroop"
	m1.YearOfRelease = 1997
	m1.Actors = []string{"A", "B"}
	m1.Grade = []int{1, 2, 3}

	i := 0
	for _, grade := range m1.Grade {
		//fmt.Println(grade)
		i = i + grade
		//fmt.Println(i)
	}
	m1.AvgOfGrade = i / len(m1.Grade)

	fmt.Println(m1.AvgOfGrade)
}
