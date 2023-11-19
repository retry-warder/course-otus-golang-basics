package types

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"
)

type Data struct {
	listsignals [10]int
	sum         float64
	avg         float64
}

func GenSignal(signal chan<- int, t time.Duration) {
	defer close(signal)
	tt := time.NewTimer(t)
	max := big.NewInt(1000)
	for {
		select {
		case <-tt.C:
			{
				return
			}
		default:
			{
				time.Sleep(10 * time.Millisecond)
				i, e := rand.Int(rand.Reader, max)
				if e != nil {
					fmt.Print(e)
				}
				signal <- int(i.Int64())
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
		d.listsignals[i-1] = s
		sum += float64(s)
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
