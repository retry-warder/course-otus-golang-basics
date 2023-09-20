package main

import (
	"fmt"

	"github.com/retry-warder/course-otus-golang-basics/hw07_word_counter/wordcounter"
)

func main() {
	inText := "Умереть за любовь не сложно. Сложно найти любовь, за которую стоит умереть"
	mapWC := wordcounter.CountWord(inText)
	fmt.Println(mapWC)
	fmt.Println(len(mapWC))
}
