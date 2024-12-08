package main

import (
	"math/rand"
	"time"
)

func RandomNumberGenerator(count int, ch chan<- int) {
	defer close(ch)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < count; i++ {
		ch <- r.Intn(100)
	}
}

func main() {
	ch := make(chan int)
	go RandomNumberGenerator(10, ch)

	for num := range ch {
		println(num)
	}
}
