package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

func multiply(v int) int {
	return v * 2
}

func generate() chan int {
	data := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(250 * time.Millisecond)
			data <- i
		}
		close(data)
	}()
	return data
}

func doWork(c <-chan int, id int8) chan int {
	out := make(chan int)
	go func(id int8) {
		for i := range c {
			fmt.Printf("Channel %v: reading %v\n", id, i)
			out <- multiply(i)
			time.Sleep(time.Second)
		}
		close(out)
	}(id)
	return out
}

func fanIn[T any](chans ...<-chan T) <-chan T {
	out := make(chan T)
	wg := &sync.WaitGroup{}
	wg.Add(len(chans))
	for _, c := range chans {
		go func(c <-chan T) {
			for r := range c {
				out <- r
			}
			wg.Done()
		}(c)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func Range(bottom, top float64, opts ...float64) []float64 {
	if top <= bottom {
		return make([]float64, 0)
	}
	diff := top - bottom
	var step float64
	if len(opts) > 0 {
		step = opts[0]
	} else {
		step = 1
	}
	length := int(math.Floor(diff / step))
	result := make([]float64, length)
	for i, v := 0, bottom; i < length; i++ {
		v += step
		result[i] = v
	}
	return result
}

func main() {

	ranges := [][]float64{
		Range(0, 10, 1),
		Range(0, 10),
		Range(0, 10, 2),
		Range(0, 5, 0.5),
		Range(5, 10, 1),
		Range(10, 1, 1),
		Range(0, 1.1, 0.3),
		Range(1, 1.1, 0.3),
	}
	for _, r := range ranges {
		fmt.Println(r, len(r), cap(r))
	}

}
