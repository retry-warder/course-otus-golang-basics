package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Data struct {
	signals [10]int
	sum     float64
	avg     float64
}

func GenSignal(signal chan<- int) {
	defer close(signal)
	tt := time.NewTimer(2 * time.Second)
	for {
		select {
		case <-tt.C:
			{
				return
			}
		default:
			{
				time.Sleep(10 * time.Millisecond)
				i := rand.Intn(1000)
				signal <- i
			}
		}
	}
}

func DataProcess(signal <-chan int, data chan<- Data) {
	var i, s int
	var sum float64
	var d Data
	defer close(data)
	for s = range signal {
		i++
		d.signals[i-1] = s
		sum = sum + float64(s)
		if i == 10 {
			d.sum = sum
			d.avg = sum / 10
			data <- d
			i = 0
			sum = 0
			d = Data{}
		}

	}
}

func main() {
	signal := make(chan int)
	data := make(chan Data)
	go GenSignal(signal)
	go DataProcess(signal, data)
	for d := range data {
		fmt.Println("------------------------------")
		fmt.Println(d)
	}
}
