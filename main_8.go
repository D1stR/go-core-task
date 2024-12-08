package main

import (
	"fmt"
	"time"
)

type SemaphoreWaitGroup struct {
	count int
	sem   chan struct{}
}

func NewSemaphoreWaitGroup() *SemaphoreWaitGroup {
	return &SemaphoreWaitGroup{
		sem: make(chan struct{}, 1),
	}
}

func (wg *SemaphoreWaitGroup) Add(delta int) {
	wg.sem <- struct{}{}
	wg.count += delta
}

func (wg *SemaphoreWaitGroup) Done() {
	wg.count--
	<-wg.sem
}

func (wg *SemaphoreWaitGroup) Wait() {
	for {
		if wg.count == 0 {
			return
		}
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	wg := NewSemaphoreWaitGroup()

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Printf("Горутина %d завершена\n", i)
			time.Sleep(1 * time.Second)
		}(i)
	}

	wg.Wait()
	fmt.Println("Все горутины завершены")
}
