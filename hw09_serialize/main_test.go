package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
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

	var b3 []types.Book
	b3 = append(b3, b1, b2)

	f, err := os.Open("test.json")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	j3, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}

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
			j, er := json.Marshal(tc.input)
			require.NoError(t, er)
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

	var b3 []types.Book
	b3 = append(b3, b1, b2)
	f, _ := os.Open("test.json")
	defer f.Close()
	j3, _ := io.ReadAll(f)

	strtests := []struct {
		name     string
		input    []byte
		expected string
	}{
		{"Ok Unmarshall json 1", j1, fmt.Sprintf("%v", b1)},
		{"Ok Unmarshall json 2", j2, fmt.Sprintf("%v", b2)},
		{"Ok Unmarshall slice json", j3, fmt.Sprintf("%v", b3)},
	}

	for _, tc := range strtests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			if tc.name == "Ok Unmarshall slice json" {
				var res []types.Book
				err := json.Unmarshal(tc.input, &res)
				require.NoError(t, err)
				require.Equal(t, tc.expected, fmt.Sprintf("%v", res),
					fmt.Sprintf("input (%v) = %v; want %v", tc.input, res, tc.expected))
			} else {
				var res types.Book
				err := json.Unmarshal(tc.input, &res)
				require.NoError(t, err)
				require.Equal(t, tc.expected, fmt.Sprintf("%v", res),
					fmt.Sprintf("input (%v) = %v; want %v", tc.input, res, tc.expected))
			}
		})
	}
}

func Test_MarshallPB(t *testing.T) {
	strtests := []struct {
		name     string
		input    types.Book
		expected int
	}{
		{"Ok marshall pb", types.NewBook(1, "Test 1", "Test.T.T", 2010, 1000, 10.10), 31},
		{"Ok marshall pb", types.NewBook(1, "Test 2", "Тест Т.Т.", 2000, 1500, 10.90), 38},
	}
	for _, tc := range strtests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			b := types_pb.NewBook(uint32(tc.input.ID),
				tc.input.Title,
				tc.input.Author,
				uint32(tc.input.Year),
				uint32(tc.input.Size),
				float32(tc.input.Rate))
			m, err := proto.Marshal(b)
			require.NoError(t, err)
			res := len(m)
			require.Equal(t, tc.expected, res, fmt.Sprintf("input (%v) = %v; want %v", b.String(), res, tc.expected))
		})
	}
}
