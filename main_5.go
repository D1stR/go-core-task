package main

import (
	"fmt"
)

func Intersection(a, b []int) (bool, []int) {
	intersectMap := make(map[int]bool)
	result := make([]int, 0)

	for _, v := range a {
		intersectMap[v] = true
	}

	for _, v := range b {
		if intersectMap[v] {
			result = append(result, v)
		}
	}

	return len(result) > 0, result
}

func main() {
	a := []int{65, 3, 58, 678, 64}
	b := []int{64, 2, 3, 43}

	hasIntersection, intersectionValues := Intersection(a, b)
	fmt.Printf("Есть ли пересечение? %t\n", hasIntersection)
	fmt.Println("Пересекающиеся значения:", intersectionValues)
}
