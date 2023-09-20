package wordcounter

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_CountWord_Text(t *testing.T) {
	ResOk := "map[Сложно:1 Умереть:1 за:2 которую:1 любовь:2 найти:1 не:1 сложно:1 стоит:1 умереть:1]"
	ResEmpty := "map[]"
	tests := []struct {
		name     string
		inText   string
		expected string
	}{
		{"empty", "", ResEmpty},
		{"ok", "Умереть за любовь не сложно. Сложно найти любовь, за которую стоит умереть", ResOk},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			res := fmt.Sprint(CountWord(tc.inText))
			require.Equal(t, tc.expected, res, fmt.Sprintf("get = %v; want %v", res, tc.expected))
		})
	}
}

func Test_CountWord_Map(t *testing.T) {
	MapOk := make(WordCountMap)
	MapOk["Сложно"] = 1
	MapOk["Умереть"] = 1
	MapOk["за"] = 2
	MapOk["которую"] = 1
	MapOk["любовь"] = 2
	MapOk["найти"] = 1
	MapOk["не"] = 1
	MapOk["сложно"] = 1
	MapOk["стоит"] = 1
	MapOk["умереть"] = 1
	MapEmpty := make(WordCountMap)
	tests := []struct {
		name     string
		inText   string
		expected WordCountMap
	}{
		{"empty", "", MapEmpty},
		{"ok", "Умереть за любовь не сложно. Сложно найти любовь, за которую стоит умереть", MapOk},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			res := CountWord(tc.inText)
			require.Equal(t, tc.expected, res, fmt.Sprintf("get = %v; want %v", res, tc.expected))
		})
	}
}

func Test_CoontWord_Len(t *testing.T) {
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
			res := len(CountWord(tc.inText))
			require.Equal(t, tc.expected, res, fmt.Sprintf("get = %d; want %v", res, tc.expected))
		})
	}
}
