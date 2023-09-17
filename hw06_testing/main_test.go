package hw06_testing

import "testing"

func Test_HW03(t *testing.T) {
	tests := []struct {
		name     string
		size     int
		expected string
	}{
		{"empty", 0, ""},
		{"ok", 3, ""},
	}
}
