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

	var number8 int8
	number8 = 127
	fmt.Println(number8)

	var number16 int16
	number16 = 567
	fmt.Println(number16)

	result := int16(number8) + number16
	fmt.Println(result)
}
