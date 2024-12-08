package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	map1 := NewStringIntMap()
	map1.Add("Unicorns", 100)
	v, ok := map1.Get("Unicorns")
	if !ok || v != 100 {
		t.Errorf("Ожидаемое значение у  'Unicorns' должно быть 100, получено %d", v)
	}
}

func TestRemove(t *testing.T) {
	map1 := NewStringIntMap()
	map1.Add("ApplicationServer", 200)
	map1.Remove("ApplicationServer")
	_, ok := map1.Get("ApplicationServer")
	if ok {
		t.Error("Ключ 'ApplicationServer' не должен существовать после удаления")
	}
}

func TestCopy(t *testing.T) {
	map1 := NewStringIntMap()
	map1.Add("ApplicationServer", 300)
	map1.Add("WebServer", 400)
	copyMap := map1.Copy()
	if copyMap["ApplicationServer"] != 300 || copyMap["WebServer"] != 400 {
		t.Errorf("Скопированная карта не соответствует оригинальной карте")
	}
}

func TestExists(t *testing.T) {
	map1 := NewStringIntMap()
	map1.Add("ApplicationServer", 500)
	if !map1.Exists("ApplicationServer") {
		t.Error("Ключ 'ApplicationServer' должен быть в мапе")
	}
	if map1.Exists("nonexistent_key") {
		t.Error("Nonexistent key должен возвращать false")
	}
}

func TestGet(t *testing.T) {
	map1 := NewStringIntMap()
	map1.Add("ApplicationServer", 600)
	v, ok := map1.Get("ApplicationServer")
	if !ok || v != 600 {
		t.Errorf("Ожидаемое значение у 'ApplicationServer' должно быть 600, получено %d", v)
	}
	_, ok = map1.Get("nonexistent_key")
	if ok {
		t.Error("Получение nonexistent key должно возвращать false")
	}
}
