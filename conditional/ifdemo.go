package conditional

import "fmt"

func IfDemo() {

	batsMen := 101

	if batsMen >= 100 {
		fmt.Println("Batsmen scored a century")
	} else {
		fmt.Println("Batsmen didn't score a century")
	}

	if true {
		fmt.Println("statemnet 1")
	} else {
		fmt.Println("statement ")
	}

	if !(10 == 15) {
		fmt.Println("a")

	} else {
		fmt.Println("b")
	}

	a := 80
	if a >= 100 && a <= 200 {
		fmt.Println("century")
	} else if a >= 200 {
		fmt.Println("double century")
	} else {
		fmt.Println("some runs")
	}

	ab := "sage"

	switch ab {
	case "sage":
		fmt.Println("the sage")
	}

}
