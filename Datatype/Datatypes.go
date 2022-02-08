package Datatype

import "fmt"

func DataTypes() {

	var a bool
	a = true
	fmt.Println(a)
	a = false
	fmt.Println(a)

	var str string
	str = "Hello Datatypes"
	fmt.Println(str)

	var number int
	number = 45
	fmt.Println(number)
	fmt.Printf("%v,%T\n", number, number)
}
