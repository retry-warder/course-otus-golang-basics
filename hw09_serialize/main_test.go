package main

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/retry-warder/course-otus-golang-basics/hw09_serialize/types"
	"github.com/retry-warder/course-otus-golang-basics/hw09_serialize/types_pb"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

func Test_MarshallJson(t *testing.T) {
	b1 := types.NewBook(1, "Test 1", "Test.T.T", 2010, 1000, 10.10)
	j1 := []byte(`{"id":1,"title":"Test 1","author":"Test.T.T","year":2010,"size":1000,"rate":10.1}`)

	b2 := types.NewBook(1, "Test 2", "Test.X.Y", 2000, 1500, 10.90)
	j2 := []byte(`{"id":1,"title":"Test 2","author":"Test.X.Y","year":2000,"size":1500,"rate":10.9}`)

	var b3 types.Books
	b3 = append(b3, b1, b2)
	j3 := []byte(`[{"id":1,"title":"Test 1","author":"Test.T.T","year":2010,"size":1000,"rate":10.1},{"id":1,"title":"Test 2","author":"Test.X.Y","year":2000,"size":1500,"rate":10.9}]`)

	strtests := []struct {
		name     string
		input    any
		expected []byte
	}{
		{"Ok marshall json", b1, j1},
		{"Ok marshall json", b2, j2},
		{"Ok marshall slice json", b3, j3},
	}

	for _, tc := range strtests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			j, err := json.Marshal(tc.input)
			require.NoError(t, err)
			res := j
			require.Equal(t, tc.expected, res, fmt.Sprintf("input (%v) = %v; want %v", tc.input, res, tc.expected))
		})

	}
}

func Test_UnmarshallJson(t *testing.T) {
	b1 := types.NewBook(1, "Test 1", "Test.T.T", 2010, 1000, 10.10)
	j1 := []byte(`{"id":1,"title":"Test 1","author":"Test.T.T","year":2010,"size":1000,"rate":10.1}`)

	b2 := types.NewBook(1, "Test 2", "Test.X.Y", 2000, 1500, 10.90)
	j2 := []byte(`{"id":1,"title":"Test 2","author":"Test.X.Y","year":2000,"size":1500,"rate":10.9}`)

	var b3 types.Books
	b3 = append(b3, b1, b2)
	j3 := []byte(`[{"id":1,"title":"Test 1","author":"Test.T.T","year":2010,"size":1000,"rate":10.1},{"id":1,"title":"Test 2","author":"Test.X.Y","year":2000,"size":1500,"rate":10.9}]`)

	strtests := []struct {
		name     string
		input    []byte
		expected string
	}{
		{"Ok Unmarshall json", j1, fmt.Sprintf("%v", b1)},
		{"Ok Unmarshall json", j2, fmt.Sprintf("%v", b2)},
		{"Ok Unmarshall slice json", j3, fmt.Sprintf("%v", b3)},
	}

	for _, tc := range strtests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			if tc.name == "Ok Unmarshall slice json" {
				var res types.Books
				err := json.Unmarshal(tc.input, &res)
				require.NoError(t, err)
				require.Equal(t, tc.expected, fmt.Sprintf("%v", res), fmt.Sprintf("input (%v) = %v; want %v", tc.input, res, tc.expected))
			} else {
				var res types.Book
				err := json.Unmarshal(tc.input, &res)
				require.NoError(t, err)
				require.Equal(t, tc.expected, fmt.Sprintf("%v", res), fmt.Sprintf("input (%v) = %v; want %v", tc.input, res, tc.expected))
			}

		})

	}
}

func Test_MarshallPB(t *testing.T) {
	strtests := []struct {
		name     string
		input    types_pb.Book
		expected int
	}{
		{"Ok marshall pb", *types_pb.NewBook(1, "Test 1", "Test.T.T", 2010, 1000, 10.10), 31},
		{"Ok marshall pb", *types_pb.NewBook(1, "Test 2", "Тест Т.Т.", 2000, 1500, 10.90), 38},
	}

	for _, tc := range strtests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			m, err := proto.Marshal(&tc.input)
			require.NoError(t, err)
			res := len(m)
			require.Equal(t, tc.expected, res, fmt.Sprintf("input (%v) = %v; want %v", tc.input.String(), res, tc.expected))
		})

	}
}

func Test_MarshallSlicePB(t *testing.T) {

	var books types_pb.Books
	b1 := types_pb.NewBook(1, "Test 1", "Test.T.T", 2010, 1000, 10.10)
	b2 := types_pb.NewBook(1, "Test 2", "Тест Т.Т.", 2000, 1500, 10.90)
	var listbooks []*types_pb.Book
	listbooks = append(listbooks, b1, b2)
	books.Value = listbooks

	l3 := 73

	strtests := []struct {
		name     string
		input    types_pb.Books
		expected int
	}{
		{"Ok marshall slice pb", books, l3},
	}

	for _, tc := range strtests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			m, err := proto.Marshal(&tc.input)
			require.NoError(t, err)
			res := len(m)
			require.Equal(t, tc.expected, res, fmt.Sprintf("input (%v) = %v; want %v", tc.input.String(), res, tc.expected))
		})

	}
}

func Test_UnmarshallPB(t *testing.T) {
	b1 := types_pb.NewBook(1, "Test 1", "Test.T.T", 2010, 1000, 10.10)
	b2 := types_pb.NewBook(1, "Test 2", "Тест Т.Т.", 2000, 1500, 10.90)

	s1 := b1.String()
	s2 := b2.String()

	m1, _ := proto.Marshal(b1)
	m2, _ := proto.Marshal(b2)

	strtests := []struct {
		name     string
		input    []byte
		expected string
	}{
		{"Ok Unmarshall pb", m1, s1},
		{"Ok Unmarshall pb", m2, s2},
	}

	for _, tc := range strtests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			var res types_pb.Book
			err := proto.Unmarshal(tc.input, &res)
			require.NoError(t, err)
			require.Equal(t, tc.expected, res.String(), fmt.Sprintf("input (%v) = %v; want %v", tc.input, res.String(), tc.expected))
		})

	}
}

func Test_UnmarshallPBSlice(t *testing.T) {
	var books types_pb.Books
	b1 := types_pb.NewBook(1, "Test 1", "Test.T.T", 2010, 1000, 10.10)
	b2 := types_pb.NewBook(1, "Test 2", "Тест Т.Т.", 2000, 1500, 10.90)
	var listbooks []*types_pb.Book
	listbooks = append(listbooks, b1, b2)
	books.Value = listbooks

	s1 := books.String()

	m, _ := proto.Marshal(&books)

	strtests := []struct {
		name     string
		input    []byte
		expected string
	}{
		{"Ok Unmarshall pb", m, s1},
	}

	for _, tc := range strtests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			var res types_pb.Books
			err := proto.Unmarshal(tc.input, &res)
			require.NoError(t, err)
			require.Equal(t, tc.expected, res.String(), fmt.Sprintf("input (%v) = %v; want %v", tc.input, res.String(), tc.expected))
		})

	}
}
