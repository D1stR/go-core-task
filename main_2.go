package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateOriginalSlice() []int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var originalSlice = make([]int, 10)
	for i := range originalSlice {
		originalSlice[i] = r.Intn(100)
	}
	return originalSlice
}

func sliceExample(slice []int) []int {
	var result []int
	for _, v := range slice {
		if v%2 == 0 {
			result = append(result, v)
		}
	}
	return result
}

func addElements(slice []int, element int) []int {
	newSlice := make([]int, len(slice), len(slice)+1)
	copy(newSlice, slice)
	newSlice = append(newSlice, element)
	return newSlice
}

func copySlice(slice []int) []int {
	copiedSlice := make([]int, len(slice))
	copy(copiedSlice, slice)
	return copiedSlice
}

func removeElement(slice []int, index int) []int {
	if index < 0 || index >= len(slice) {
		return slice
	}
	return append(slice[:index], slice[index+1:]...)
}

func main() {
	originalSlice := generateOriginalSlice()
	fmt.Println("Оригинальный слайс:", originalSlice)

	evenNumbers := sliceExample(originalSlice)
	fmt.Println("Четные числа:", evenNumbers)

	appendedSlice := addElements(originalSlice, 42)
	fmt.Println("Слайс после добавления элемента:", appendedSlice)

	copiedSlice := copySlice(originalSlice)
	fmt.Println("Копия слайса:", copiedSlice)

	removedIndex := 3
	removedSlice := removeElement(originalSlice, removedIndex)
	fmt.Printf("Слайс после удаления элемента по индексу %d: %v\n", removedIndex, removedSlice)

}
