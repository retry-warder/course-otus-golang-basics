package binarysearch

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_BinarySearch_1(t *testing.T) {
	testval := []int{1, 3, 7, 9, 11, 31, 56, 75, 116}
	tests := []struct {
		name     string
		value    int
		expected bool
	}{
		{"ok", 11, true},
		{"ok", 31, true},
		{"fail", 10, false},
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
	testval := []int{1, 3, 7, 9, 11, 31, 56, 75, 116}
	tests := []struct {
		name     string
		value    int
		expected int
	}{
		{"ok", 11, 4},
		{"ok", 31, 5},
		{"fail", 10, -1},
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
