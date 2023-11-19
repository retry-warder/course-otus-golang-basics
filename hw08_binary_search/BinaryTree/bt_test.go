package binarytree

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_BSearch_1(t *testing.T) {
	testval := []int{8, 4, 2, 3, 10, 6, 7, 32, 16}
	bt := NewBst()
	for _, i := range testval {
		bt.Insert(i, fmt.Sprintf("Bingo_%d", i))
	}
	tests := []struct {
		name     string
		value    int
		expected bool
	}{
		{"ok", 4, true},
		{"ok", 16, true},
		{"fail", 11, false},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			res, _ := BSearch(bt.root, tc.value)
			require.Equal(t, tc.expected, res, fmt.Sprintf("get = %v; want %v", res, tc.expected))
		})
	}
}

func Test_BSearch_2(t *testing.T) {
	testval := []int{8, 4, 2, 3, 10, 6, 7, 32, 16}
	bt := NewBst()
	for _, i := range testval {
		bt.Insert(i, fmt.Sprintf("Bingo_%d", i))
	}
	tests := []struct {
		name     string
		value    int
		expected string
	}{
		{"ok", 4, "Bingo_4"},
		{"ok", 16, "Bingo_16"},
		{"fail", 11, ""},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			_, res := BSearch(bt.root, tc.value)
			require.Equal(t, tc.expected, res, fmt.Sprintf("get = %v; want %v", res, tc.expected))
		})
	}
}
