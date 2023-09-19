package main

import (
	"fmt"

	"github.com/retry-warder/course-otus-golang-basics/hw05_shapes/hw07_word_counter/types"
)

func main() {
	inText := "Умереть за любовь не сложно. Сложно найти любовь, за которую стоит умереть"
	mapWC := types.CountWord(inText)
	fmt.Println(mapWC)
	fmt.Println(len(mapWC))
}
