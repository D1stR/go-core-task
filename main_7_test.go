package main

import (
	"testing"
)

func TestMergeChannels(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		ch1 <- 1
		ch1 <- 2
		close(ch1)
	}()

	go func() {
		ch2 <- 3
		ch2 <- 4
		close(ch2)
	}()

	merged := merge(ch1, ch2)

	expected := []int{1, 2, 3, 4}
	var result []int

	for val := range merged {
		result = append(result, val)
	}

	if len(result) != len(expected) {
		t.Errorf("Ожидаемое значение %d, получено %d", len(expected), len(result))
	}

	for i, v := range result {
		if v != expected[i] {
			t.Errorf("Ожидали %d, получили %d", expected[i], v)
		}
	}
}
