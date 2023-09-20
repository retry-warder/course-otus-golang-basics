package wordcounter

import "strings"

type WordCountMap map[string]int64

const WordCountSep string = " \r\n\t,.:;!?\"'()[]"

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
				(mapWC)[word]++
				buffer.Reset()
			}
		}
	}
	delete((mapWC), "-")
	return mapWC
}
