package main

import (
	"fmt"
)

type StringIntMap struct {
	data map[string]int
}

func NewStringIntMap() *StringIntMap {
	return &StringIntMap{
		data: make(map[string]int),
	}
}

func (m *StringIntMap) Add(key string, value int) {
	m.data[key] = value
}

func (m *StringIntMap) Remove(key string) {
	delete(m.data, key)
}

func (m *StringIntMap) Copy() map[string]int {
	copied := make(map[string]int, len(m.data))
	for k, v := range m.data {
		copied[k] = v
	}
	return copied
}

func (m *StringIntMap) Exists(key string) bool {
	_, ok := m.data[key]
	return ok
}

func (m *StringIntMap) Get(key string) (int, bool) {
	value, exists := m.data[key]
	return value, exists
}

func main() {
	map1 := NewStringIntMap()
	map1.Add("GOLANG", 10)
	map1.Add("Docker", 20)
	map1.Add("Astana", 450)
	map1.Add("Audi", 9999)
	fmt.Println(map1.Get("Audi"))
	fmt.Println(map1.Exists("Docker"))
	fmt.Println(map1.Copy())
	map1.Remove("Docker")
	fmt.Println(map1.Get("GOLANG"))
}
