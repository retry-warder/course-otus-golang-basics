package types

import "strings"

type WordCountMap map[string]int64

const WordCountSep string = " \n\t,.:;!?\"'()[]"

func CountWord(inText string) WordCountMap {
	var buffer strings.Builder

	mapWC := make(WordCountMap)

	mapSep := make(map[rune]struct{})
	for _, symbol := range WordCountSep {
		mapSep[symbol] = struct{}{}
	}

	arrText := []rune(inText)
	for i, v := range arrText {
		_, oks := mapSep[v]
		if !oks {
			buffer.WriteRune(v)
		}
		if oks || i == len(arrText)-1 {
			if buffer.Len() > 0 {
				word := buffer.String()
				_, okw := (mapWC)[word]
				if okw {
					(mapWC)[word]++
				} else {
					(mapWC)[word] = 1
				}
				buffer.Reset()
			}
		}
	}
	delete((mapWC), "-")
	return mapWC
}
