package looping

import "fmt"

func LoopingDemo() {
	Grif := make([]int, 0, 40)

	for i := 0; i <= 10; i++ {
		Grif = append(Grif, i)
		fmt.Println(Grif[i])

	}

	loopvariable := 0
	for loopvariable <= 5 {
		fmt.Println(loopvariable)
		loopvariable++
	}

}
