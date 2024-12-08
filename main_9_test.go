package main

import (
	"testing"
)

func TestGenerateNumbers(t *testing.T) {
	inChan := make(chan uint8)

	go generateNumbers(inChan)

	expectedCount := 100
	count := 0

	for num := range inChan {
		if num != uint8(count) {
			t.Errorf("Ожидали %d, получили %d", count, num)
		}
		count++
	}

	if count != expectedCount {
		t.Errorf("Ожидали %d чисел, получили %d", expectedCount, count)
	}
}

func TestCubeAndSend(t *testing.T) {
	inChan := make(chan uint8)
	outChan := make(chan NumberCube)

	go func() {
		for i := uint8(0); i < 10; i++ {
			inChan <- i
		}
		close(inChan)
	}()

	go cubeAndSend(inChan, outChan)

	expectedCubes := map[uint8]float64{
		0: 0,
		1: 1,
		2: 8,
		3: 27,
		4: 64,
		5: 125,
		6: 216,
		7: 343,
		8: 512,
		9: 729,
	}

	for result := range outChan {
		expectedCube, exists := expectedCubes[result.Number]
		if !exists {
			t.Errorf("Неожиданное число: %d", result.Number)
			continue
		}
		if result.Cube != expectedCube {
			t.Errorf("Ожидаемый куб от числа %d должен быть %.2f, получили число %.2f", result.Number, expectedCube, result.Cube)
		}
	}
}
