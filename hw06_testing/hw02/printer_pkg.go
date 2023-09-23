package hw02

import (
	"fmt"
)

func PrintStaff(staff []Employee) {
	var str string
	for i := 0; i < len(staff); i++ {
		fmt.Println(staff[i])
	}
	fmt.Println(str)
}
