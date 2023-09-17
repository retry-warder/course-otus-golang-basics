package hw02_printer

import (
	"fmt"

	hw02_types "github.com/retry-warder/course-otus-golang-basics/hw06_testing/hw02/types"
)

func PrintStaff(staff []hw02_types.Employee) {
	var str string
	for i := 0; i < len(staff); i++ {
		fmt.Println(staff[i])
	}
	fmt.Println(str)
}
