package binarysearch

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_BinarySearch_1(t *testing.T) {
	testval := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	tests := []struct {
		name     string
		value    int
		expected bool
	}{
		{"ok", 1, true},
		{"ok", 7, true},
		{"fail", 14, false},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			res, _ := BinarySearch(testval, tc.value)
			require.Equal(t, tc.expected, res, fmt.Sprintf("get = %v; want %v", res, tc.expected))
		})
	}
}

func Test_BinarySearch_2(t *testing.T) {
	testval := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	tests := []struct {
		name     string
		value    int
		expected int
	}{
		{"ok", 1, 1},
		{"ok", 8, 8},
		{"fail", 17, -1},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			_, res := BinarySearch(testval, tc.value)
			require.Equal(t, tc.expected, res, fmt.Sprintf("get = %v; want %v", res, tc.expected))
		})
	}
}
