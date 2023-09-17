package hw06_testing

import (
	"fmt"
	"math"
	"testing"

	hw02_reader "github.com/retry-warder/course-otus-golang-basics/hw06_testing/hw02/reader"
	hw02_types "github.com/retry-warder/course-otus-golang-basics/hw06_testing/hw02/types"
	"github.com/retry-warder/course-otus-golang-basics/hw06_testing/hw03"
	hw04_types "github.com/retry-warder/course-otus-golang-basics/hw06_testing/hw04/types"
	hw05_types "github.com/retry-warder/course-otus-golang-basics/hw06_testing/hw05/types"
	"github.com/stretchr/testify/require"
)

func Test_HW03(t *testing.T) {
	str_ok := "# # # # \r\n # # # #\r\n# # # # \r\n # # # #\r\n# # # # \r\n # # # #\r\n# # # # \r\n # # # #\r\n"
	str_fail := "# # #\r\n # # # #\r\n# # # # \r\n # # # #\r\n# # # # \r\n # # # #\r\n## # \r\n # # # #\r\n"
	tests := []struct {
		name     string
		size     int
		expected string
	}{
		{"empty", 0, ""},
		{"ok", 8, str_ok},
		{"fail", 8, str_fail},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			res := hw03.Chessboard(&tc.size)
			require.Equal(t, tc.expected, res, fmt.Sprintf("Chessboard(%d) = %v; want %v", tc.size, res, tc.expected))

		})
	}
}
func Test_HW04(t *testing.T) {
	var res string
	book1 := hw04_types.NewBook(1, "T3", "A1", 2022, 100, 10)
	book2 := hw04_types.NewBook(2, "T2", "A2", 2023, 120, 8)

	tests := []struct {
		name         string
		book1        *hw04_types.Book
		book2        *hw04_types.Book
		type_compare hw04_types.TypeCompare
		expected     string
	}{
		{"Ok Compare YEAR", book1, book2, hw04_types.Year, book2.String()},
		{"Ok Compare SIZE", book1, book2, hw04_types.Size, book2.String()},
		{"Ok Compare RATE", book1, book2, hw04_types.Rate, book1.String()},
		{"FAIL Compare YEAR", book1, book2, hw04_types.Year, book1.String()},
		{"FAIL Compare SIZE", book1, book2, hw04_types.Size, book1.String()},
		{"FAIL Compare RATE", book1, book2, hw04_types.Rate, book2.String()},
	}
	for _, tc := range tests {
		compare := hw04_types.NewCompareBook(tc.type_compare)
		if compare.Compare(tc.book1, tc.book2) {
			res = tc.book1.String()
		} else {
			res = tc.book2.String()
		}
		if res != tc.expected {
			t.Errorf("Compare book by (%v) = %v; want %v", tc.type_compare, res, tc.expected)
		}

	}
}

func Test_HW05_Ok(t *testing.T) {
	tests := []struct {
		name     string
		shape    hw05_types.Shape
		expected float64
		err      error
	}{
		{"Ok Calc Area of Circle", hw05_types.NewCircle(8), float64(math.Pi) * math.Pow(float64(8), 2), nil},
		{"Ok Calc Area of Rectangle", hw05_types.NewRectangle(2, 5), float64(2) * float64(5), nil},
		{"Ok Calc Area of Triangle", hw05_types.NewTriangle(5, 10), (float64(5) * float64(10)) / 2, nil},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			area, _ := hw05_types.CalculateArea(tc.shape)
			require.Equal(t, tc.expected, area, fmt.Sprintf("Calc Area (%v) = %f; want %f", tc.shape, area, tc.expected))
		})
	}
}

func Test_HW05_Fail(t *testing.T) {
	area, err := hw05_types.CalculateArea("shape")
	if area == 0.0 && err != nil {
		t.Errorf("Fail Calc Area (%v)", err)
	}
}

func Test_HW02(t *testing.T) {
	var err error
	var staff []hw02_types.Employee
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Ok", "hw02/test_1.json", fmt.Sprintf("User ID: %d; Age: %d; Name: %s; Department ID: %d; ", 1, 22, "Ivan", 2)},
		{"Fail", "hw02/test_2.json", fmt.Sprintf("User ID: %d; Age: %d; Name: %s; Department ID: %d; ", 1, 27, "Eva", 2)},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			staff, err = hw02_reader.ReadJSON(tc.input)
			require.NoError(t, err)
			res := staff[0].String()
			require.Equal(t, tc.expected, res, fmt.Sprintf("input (%v) = %v; want %v", tc.input, res, tc.expected))
		})
	}
}
