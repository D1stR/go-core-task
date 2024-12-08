package main

import (
	"reflect"
	"testing"
)

func TestDifference(t *testing.T) {
	tests := []struct {
		name   string
		slice1 []string
		slice2 []string
		want   []string
	}{
		{
			name:   "Обычный пример",
			slice1: []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"},
			slice2: []string{"banana", "date", "fig"},
			want:   []string{"apple", "cherry", "43", "lead", "gno1"},
		},
		{
			name:   "Второй слайс пустой",
			slice1: []string{"apple", "banana", "cherry"},
			slice2: []string{},
			want:   []string{"apple", "banana", "cherry"},
		},
		{
			name:   "Нет общих элементов",
			slice1: []string{"apple", "banana", "cherry"},
			slice2: []string{"date", "fig", "grape"},
			want:   []string{"apple", "banana", "cherry"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := difference(tt.slice1, tt.slice2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("difference() = %v, want %v", got, tt.want)
			}
		})
	}
}
