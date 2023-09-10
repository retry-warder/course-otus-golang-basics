package main

import (
	"fmt"

	"github.com/retry-warder/course-otus-golang-basics/hw02_fix_app/printer"
	"github.com/retry-warder/course-otus-golang-basics/hw02_fix_app/reader"
	"github.com/retry-warder/course-otus-golang-basics/hw02_fix_app/types"
)

func main() {
	var path string

	fmt.Printf("Enter data file path: ")
	fmt.Scanln(&path)

	var err error
	var staff []types.Employee

	if len(path) == 0 {
		path = "data.json"
	}

	staff, err = reader.ReadJSON(path)
	if err != nil {
		fmt.Print(err)
	}

	printer.PrintStaff(staff)
}
