package main

import (
	"testing"
)

func TestGenerateOriginalSlice(t *testing.T) {
	slice := generateOriginalSlice()
	if len(slice) != 10 {
		t.Errorf("Ожидаем слайс длиной 10, получили %d", len(slice))
	}
	for _, v := range slice {
		if v < 0 || v >= 100 {
			t.Errorf("Значение %d вышло за пределы [0, 100)", v)
		}
	}
}

func TestSliceExample(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6}
	expected := []int{2, 4, 6}
	result := sliceExample(input)

	if len(result) != len(expected) {
		t.Errorf("Ожидаемая длина %d, получили %d", len(expected), len(result))
	}

	for i, v := range result {
		if v != expected[i] {
			t.Errorf("Ожидаем %d с индексом %d, получили %d", expected[i], i, v)
		}
	}
}

func TestAddElements(t *testing.T) {
	slice := []int{1, 2, 3}
	element := 4
	expected := []int{1, 2, 3, 4}
	result := addElements(slice, element)

	if len(result) != len(expected) {
		t.Errorf("Ожидаемая длина %d, получили %d", len(expected), len(result))
	}

	for i, v := range result {
		if v != expected[i] {
			t.Errorf("Ожидаем %d с индексом %d, получили %d", expected[i], i, v)
		}
	}
}

func TestCopySlice(t *testing.T) {
	slice := []int{1, 2, 3}
	result := copySlice(slice)

	if len(result) != len(slice) {
		t.Errorf("Ожидаемая длина %d, получили %d", len(slice), len(result))
	}

	for i, v := range result {
		if v != slice[i] {
			t.Errorf("Ожидаем %d с индексом %d, получили %d", slice[i], i, v)
		}
	}
}

func TestRemoveElement(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}

	checkSliceEquality := func(result, expected []int) {
		if len(result) != len(expected) {
			t.Errorf("Ожидаемая длина %d, получили %d", len(expected), len(result))
			return
		}
		for i, v := range result {
			if v != expected[i] {
				t.Errorf("Ожидаем %d с индексом %d, получили %d", expected[i], i, v)
			}
		}
	}

	result := removeElement(slice, 2)
	expected := []int{1, 2, 4, 5}
	checkSliceEquality(result, expected)

	result = removeElement(slice, 10)
	checkSliceEquality(result, slice)

	result = removeElement(slice, -1)
	checkSliceEquality(result, slice)
}
