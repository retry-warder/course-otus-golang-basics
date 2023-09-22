package main

import (
	"fmt"

	binarysearch "github.com/retry-warder/course-otus-golang-basics/hw08_binary_search/BinarySearch"
	binarytree "github.com/retry-warder/course-otus-golang-basics/hw08_binary_search/BinaryTree"
)

func main() {
	testval1 := []int{8, 4, 2, 3, 10, 6, 7, 32, 16}
	bt := binarytree.NewBst()
	for _, i := range testval1 {
		bt.Insert(i, fmt.Sprintf("Bingo_%d", i))
	}
	r := bt.Root()

	_, s := binarytree.BSearch(r, 16)
	fmt.Println(s)

	_, s = binarytree.BSearch(r, 8)
	fmt.Println(s)

	_, s = binarytree.BSearch(r, 4)
	fmt.Println(s)

	_, s = binarytree.BSearch(r, 10)
	fmt.Println(s)

	testval2 := []int{1, 3, 7, 9, 11, 31, 56, 75, 116}
	_, res := binarysearch.BinarySearch(testval2, 31)
	fmt.Println(res)
}
