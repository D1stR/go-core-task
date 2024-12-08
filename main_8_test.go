package main

import (
	"sync"
	"testing"
	"time"
)

func TestSemaphoreWaitGroup(t *testing.T) {
	wg := NewSemaphoreWaitGroup()
	const goroutinesCount = 10

	var mu sync.Mutex
	var completed int

	for i := 0; i < goroutinesCount; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			time.Sleep(1000 * time.Millisecond)
			mu.Lock()
			completed++
			mu.Unlock()
		}(i)
	}

	wg.Wait()

	if completed != goroutinesCount {
		t.Errorf("Ожидаемое количество завершенных горутин %d, получено %d", goroutinesCount, completed)
	}
}

func TestSemaphoreWaitGroupMultipleAdd(t *testing.T) {
	wg := NewSemaphoreWaitGroup()
	const goroutinesCount = 5

	var mu sync.Mutex
	var completed int

	for i := 0; i < goroutinesCount; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			time.Sleep(50 * time.Millisecond)
			mu.Lock()
			completed++
			mu.Unlock()
		}(i)
	}

	wg.Wait()

	if completed != goroutinesCount {
		t.Errorf("Ожидаемое количество завершенных горутин %d, получено %d", goroutinesCount, completed)
	}
}
