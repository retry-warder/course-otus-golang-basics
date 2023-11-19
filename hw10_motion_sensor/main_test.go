package main

import (
	"fmt"
	"testing"

	"github.com/retry-warder/course-otus-golang-basics/hw10_motion_sensor/types"
	"github.com/stretchr/testify/require"
)

type StructTest struct {
	name     string
	input    []int
	expected float64
}

func Test_1(t *testing.T) {
	var i, j, s int
	var sum float64
	signal := make(chan int)
	var ST []StructTest
	go types.GenSignal(signal, 2)
	i = 0
	listsignals := make([]int, 10)
	for s = range signal {
		i++
		sum += float64(s)
		if i == 10 {
			j++
			ST = append(ST, StructTest{fmt.Sprintf("Test %v", j), listsignals, sum / 10})
			i = 0
			sum = 0
			listsignals = make([]int, 10)
		}
	}
	for _, tc := range ST {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			sum = 0
			for k := 0; k < len(tc.input); k++ {
				sum += float64(tc.input[k])
			}
			res := sum / float64(len(tc.input))
			require.NoError(t, _)
			require.Equal(t, tc.expected, res, fmt.Sprintf("avg (%v) = %v; want %v", tc.input, res, tc.expected))
		})
	}
}
