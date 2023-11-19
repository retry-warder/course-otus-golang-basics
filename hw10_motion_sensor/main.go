package main

import (
	"fmt"

	"github.com/retry-warder/course-otus-golang-basics/hw10_motion_sensor/types"
)

func main() {
	signal := make(chan int)
	data := make(chan types.Data)
	go types.GenSignal(signal, 60)
	go types.DataProcess(signal, data)
	for d := range data {
		fmt.Println("------------------------------")
		fmt.Println(d)
	}
}
