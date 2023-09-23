package hw03

import (
	"strings"
)

func Chessboard(size *int) string {
	var sb strings.Builder
	for i := 0; i < *size; i++ {
		for j := 0; j < *size; j++ {
			if (i%2+j%2)%2 == 0 {
				sb.WriteString("#")
			} else {
				sb.WriteString(" ")
			}
		}
		sb.WriteString("\r\n")
	}
	return sb.String()
}
