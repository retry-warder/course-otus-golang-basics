package printer

import (
	"fmt"

	"github.com/retry-warder/course-otus-golang-basics/hw02_fix_app/types"
)

func PrintStaff(staff []types.Employee) {
	var str string
	for i := 0; i < len(staff); i++ {
		fmt.Println(staff[i])
	}
	fmt.Println(str)
}
