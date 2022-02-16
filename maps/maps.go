package maps

import "fmt"

func MapsDemo() {

	MapsOfPopulatuion := make(map[string]int)
	MapsOfPopulatuion["dallas"] = 1094851
	MapsOfPopulatuion["boston"] = 2098754
	MapsOfPopulatuion["SFO"] = 3065999
	MapsOfPopulatuion["char"] = 0

	fmt.Println(MapsOfPopulatuion)
	fmt.Println(MapsOfPopulatuion["dallas"])
	fmt.Println(len(MapsOfPopulatuion))
	MapsOfPopulatuion["dallas"] = 1099999
	fmt.Println(&MapsOfPopulatuion)
	fmt.Println(MapsOfPopulatuion["dallas"])
	fmt.Println(MapsOfPopulatuion["chicago"])

	for key, _ := range MapsOfPopulatuion {

		fmt.Println(key)
	}

	if _, ok := MapsOfPopulatuion["chicago"]; !ok {
		fmt.Println("Key is present: ", ok)
	}
	if _, ok := MapsOfPopulatuion["char"]; ok {
		fmt.Println("Key is present: ", ok)
	}
}
