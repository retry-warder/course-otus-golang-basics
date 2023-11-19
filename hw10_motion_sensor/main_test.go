package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/retry-warder/course-otus-golang-basics/hw10_motion_sensor/types"
	"github.com/stretchr/testify/require"
)

type Stest struct {
	name     string
	input    []int
	expected float64
}

func Test_AvgSignal(t *testing.T) {
	var i, j, s int
	var sum float64
	signal := make(chan int)
	var ST []Stest
	go types.GenSignal(signal, 2*time.Second)
	i = 0
	listsignals := []int{}
	for s = range signal {
		i++
		sum += float64(s)
		listsignals = append(listsignals, s)
		if i == 10 {
			j++
			if j < 5 {
				ST = append(ST, Stest{fmt.Sprintf("Test %v", j), listsignals, sum / float64(len(listsignals))})
			}
			i = 0
			sum = 0
			listsignals = []int{}
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
			require.Equal(t, tc.expected, res, fmt.Sprintf("avg (%v) = %v; want %v", tc.input, res, tc.expected))
		})
	}
}
