package main

import (
	"testing"
)

func TestRandomNumberGenerator(t *testing.T) {
	count := 10
	ch := make(chan int)

	go RandomNumberGenerator(count, ch)

	received := 0
	for range ch {
		received++
	}

	if received != count {
		t.Errorf("Ожидаемые значения %d, полученные %d", count, received)
	}
}
