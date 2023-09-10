package printer

import (
	"fmt"

	"github.com/retry-warder/hw02_fix_app/types"
)

func PrintStaff(staff []types.Employee) {
	var str string
	for i := 0; i < len(staff); i++ {
		str := staff
		fmt.Println(str)
	}

	fmt.Println(str)
}
