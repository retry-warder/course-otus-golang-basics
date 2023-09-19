package main

import (
	"fmt"
	"testing"

	"github.com/retry-warder/course-otus-golang-basics/hw05_shapes/hw07_word_counter/types"
	"github.com/stretchr/testify/require"
)

func Test_CoontWord(t *testing.T) {
	tests := []struct {
		name     string
		inText   string
		expected int
	}{
		{"empty", "", 0},
		{"ok", "Умереть за любовь не сложно. Сложно найти любовь, за которую стоит умереть", 10},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			res := len(types.CountWord(tc.inText))
			require.Equal(t, tc.expected, res, fmt.Sprintf("get = %d; want %v", res, tc.expected))
		})
	}
}
