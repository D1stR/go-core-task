package main

import (
	"fmt"
)

func difference(slice1, slice2 []string) (result []string) {
	set := make(map[string]struct{})
	for _, item := range slice2 {
		set[item] = struct{}{}
	}

	for _, item := range slice1 {
		if _, ok := set[item]; !ok {
			result = append(result, item)
		}
	}
	return result
}

func main() {
	slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "date", "fig"}

	diff := difference(slice1, slice2)
	fmt.Printf("Разница между срезами:\n%s\n", diff)
}
