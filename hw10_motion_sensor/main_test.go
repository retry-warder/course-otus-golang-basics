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
	input    float64
	expected float64
}

func Test_AvgSignal(t *testing.T) {
	var i, j, s int
	var sum float64
	signal := make(chan int)
	var ST []Stest
	go types.GenSignal(signal, 2*time.Second)
	i = 0
	for s = range signal {
		i++
		sum += float64(s)
		if i == 10 {
			j++
			i = 0
			sum = 0
		}
	}
	for _, tc := range ST {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			res := tc.input / 10
			require.Equal(t, tc.expected, res, fmt.Sprintf("avg (%v) = %v; want %v", tc.input, res, tc.expected))
		})
	}
}

func Test_Data_Proc(t *testing.T) {
	check := false
	signal := make(chan int)
	data := make(chan types.Data)
	go types.GenSignal(signal, 2*time.Second)
	go types.DataProcess(signal, data)
	for d := range data {
		fmt.Println("------------------------------")
		fmt.Println(d)
		check = true
	}
	require.Equal(t, check, true)
}
