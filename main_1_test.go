package main

import (
	"crypto/sha256"
	"testing"
)

func TestGetTypes(t *testing.T) {
	numDecimal := 42
	numOctal := 052
	numHexadecimal := 0x2A
	pi := 3.14
	name := "Golang"
	isActive := true
	complexNum := complex(float32(1), float32(2))

	expected := []string{
		"int", "int", "int",
		"float64", "string", "bool", "complex64",
	}

	result := getTypes(numDecimal, numOctal, numHexadecimal, pi, name, isActive, complexNum)

	for i, v := range expected {
		if result[i] != v {
			t.Errorf("Ожидаем %s, получили %s", v, result[i])
		}
	}
}

func TestConvertToString(t *testing.T) {
	type args struct {
		numDecimal     int
		numOctal       int
		numHexadecimal int
		pi             float64
		name           string
		isActive       bool
		complexNum     complex64
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Correct Conversion",
			args: args{
				numDecimal:     123,
				numOctal:       0177,
				numHexadecimal: 0xFF,
				pi:             3.14159,
				name:           "Hello",
				isActive:       false,
				complexNum:     2 + 3i,
			},
			want: "123177ff3.14159Hellofalse(2+3i)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertToString(tt.args.numDecimal, tt.args.numOctal, tt.args.numHexadecimal, tt.args.pi, tt.args.name, tt.args.isActive, tt.args.complexNum); got != tt.want {
				t.Errorf("convertToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHashRuneSlice(t *testing.T) {
	runes := []rune("test")
	salt := "go-2024"
	expectedHash := sha256.Sum256([]byte("tego-2024st"))

	result := hashRuneSlice(runes, salt)

	if result != expectedHash {
		t.Errorf("Ожидаем %x, получили %x", expectedHash, result)
	}
}
