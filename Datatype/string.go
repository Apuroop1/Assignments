package Datatype

import "fmt"

func String() {

	var S1 string
	var S2 string

	S1 = "this is Apuroop"
	S2 = "Strings are immutable"

	S3 := []byte(S1)
	S4 := []byte(S2)
	S5 := string(S3) + string(S4)
	fmt.Println(S3)
	fmt.Println(S4)
	fmt.Println(S5)
}
