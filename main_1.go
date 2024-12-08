package main

import (
	"crypto/sha256"
	"fmt"
	"strconv"
)

func getTypes(val ...any) []string {
	var types []string
	for _, v := range val {
		types = append(types, fmt.Sprintf("%T", v))
	}
	return types
}

func convertToString(val ...any) string {
	var result string
	for _, v := range val {
		switch v := v.(type) {
		case int:
			result += strconv.Itoa(v)
		case float64:
			result += strconv.FormatFloat(v, 'f', -1, 64)
		case string:
			result += v
		case bool:
			result += strconv.FormatBool(v)
		case complex64:
			result += strconv.FormatComplex(complex128(v), 'f', -1, 64)
		default:
			result += fmt.Sprintf("%v", v) // Для других типов
		}
	}
	return result
}

func hashRuneSlice(runes []rune, salt string) [32]byte {
	saltedRunes := make([]rune, len(runes)+len(salt))
	copy(saltedRunes[:len(runes)/2], runes[:len(runes)/2])
	copy(saltedRunes[len(runes)/2:len(runes)/2+len(salt)], []rune(salt))
	copy(saltedRunes[len(runes)/2+len(salt):], runes[len(runes)/2:])

	hasher := sha256.New()
	hasher.Write([]byte(string(saltedRunes)))
	var result [32]byte
	copy(result[:], hasher.Sum(nil))
	return result
}

func main() {
	var numDecimal int = 42
	var numOctal int = 052
	var numHexadecimal int = 0x2A
	var pi float64 = 3.14
	var name string = "Golang"
	var isActive bool = true
	var complexNum complex64 = 1 + 2i

	fmt.Println("Типы переменных:")
	for _, t := range getTypes(numDecimal, numOctal, numHexadecimal, pi, name, isActive, complexNum) {
		fmt.Printf("\t%s\n", t)
	}

	str := convertToString(numDecimal, numOctal, numHexadecimal, pi, name, isActive, complexNum)
	fmt.Println("Объединенная строка:", str)

	runes := []rune(str)
	fmt.Println("Срез рун:", string(runes))

	hashed := hashRuneSlice(runes, "go-2024")
	fmt.Printf("Хешированный результат: %x\n", hashed)
}
