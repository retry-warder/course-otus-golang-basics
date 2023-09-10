package main

import (
	"hw04_struct_comparator/types"
)

func main() {
	book1 := types.NewBook(1, "T3", "A1", 2022, 100, 10)
	book2 := types.NewBook(2, "T2", "A2", 2023, 120, 15)
	compare := types.NewCompareBook(types.Size)
	compare.Compare(book1, book2)
}
