package array

import "fmt"

func SliceDemo() {

	//var SampleSlice []int
	//declaring a slice
	SampleSlice := make([]int, 0, 10)
	//SampleSlice[0] = 10
	//adding a new value to slice using append
	SampleSlice = append(SampleSlice, 20)
	SampleSlice = append(SampleSlice, 30)
	SampleSlice = append(SampleSlice, 40)
	fmt.Println(&SampleSlice[0])
	fmt.Println(&SampleSlice[1])
	fmt.Println(&SampleSlice[2])
	fmt.Println(SampleSlice)
	slice2 := SampleSlice[:3]
	slice3 := SampleSlice[0:2]
	slice3[0] = 100

	slice4 := SampleSlice[1:]
	SampleSlice2 := SampleSlice
	SampleSlice[2] = 110
	fmt.Println(SampleSlice2)
	fmt.Println(SampleSlice)
	fmt.Println(slice2)
	fmt.Println(slice3)
	fmt.Println(slice4)
	fmt.Println(len(SampleSlice))
	fmt.Println(cap(SampleSlice))
}
