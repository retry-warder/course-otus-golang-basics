package hw04

import hw04_types "github.com/retry-warder/course-otus-golang-basics/hw06_testing/hw04/types"

func Struct_comparator() hw04_types.Book {
	book1 := hw04_types.NewBook(1, "T3", "A1", 2022, 100, 10)
	book2 := hw04_types.NewBook(2, "T2", "A2", 2023, 120, 15)
	compare := hw04_types.NewCompareBook(hw04_types.Size)
	if compare.Compare(book1, book2) {
		return *book1
	} else {
		return *book2
	}
}
