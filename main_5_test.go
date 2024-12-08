package main

import (
	"testing"
)

func TestIntersection(t *testing.T) {
	tests := []struct {
		name       string
		sliceA     []int
		sliceB     []int
		wantHasInt bool
		wantResult []int
	}{
		{
			name:       "Пример из задания",
			sliceA:     []int{65, 3, 58, 678, 64},
			sliceB:     []int{64, 2, 3, 43},
			wantHasInt: true,
			wantResult: []int{64, 3},
		},
		{
			name:       "Нет пересечений",
			sliceA:     []int{10, 20, 30},
			sliceB:     []int{40, 50, 60},
			wantHasInt: false,
			wantResult: nil,
		},
		{
			name:       "Пустые слайсы",
			sliceA:     []int{},
			sliceB:     []int{},
			wantHasInt: false,
			wantResult: nil,
		},
		{
			name:       "Один общий элемент",
			sliceA:     []int{100, 200, 300},
			sliceB:     []int{400, 500, 600, 100},
			wantHasInt: true,
			wantResult: []int{100},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotHasInt, gotResult := Intersection(tt.sliceA, tt.sliceB)
			if gotHasInt != tt.wantHasInt {
				t.Errorf("Intersection() hasInter = %v, want %v", gotHasInt, tt.wantHasInt)
			}
			if !equalSlices(gotResult, tt.wantResult) {
				t.Errorf("Intersection() result = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func equalSlices(s1, s2 []int) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}
