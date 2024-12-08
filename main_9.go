package main

import (
	"fmt"
)

type NumberCube struct {
	Number uint8
	Cube   float64
}

func generateNumbers(inChan chan<- uint8) {
	for i := uint8(0); i < 100; i++ {
		inChan <- i
	}
	close(inChan)
}

func cubeAndSend(inChan <-chan uint8, outChan chan<- NumberCube) {
	for num := range inChan {
		cubed := float64(int(num) * int(num) * int(num))
		outChan <- NumberCube{Number: num, Cube: cubed}
	}
	close(outChan)
}

func main() {
	inChan := make(chan uint8)
	outChan := make(chan NumberCube)

	go generateNumbers(inChan)
	go cubeAndSend(inChan, outChan)

	fmt.Println("Результаты возведения в куб чисел от 0 до 99:")
	for result := range outChan {
		fmt.Printf("Число %d в кубе: %.2f\n", result.Number, result.Cube)
	}
}
