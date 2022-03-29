package main

import "testing"

func TestFreqStack(t *testing.T) {
	fs := Constructor()
	fs.Push(5)
	fs.Push(7)
	fs.Push(5)
	fs.Push(7)
	fs.Push(4)
	fs.Push(5)
	if val := fs.Pop(); val != 5 {
		t.Errorf("Pop1: want %d, got %d", 5, val)
	}
	if val := fs.Pop(); val != 7 {
		t.Errorf("Pop2: want %d, got %d", 7, val)
	}
	if val := fs.Pop(); val != 5 {
		t.Errorf("Pop3: want %d, got %d", 5, val)
	}
	if val := fs.Pop(); val != 4 {
		t.Errorf("Pop4: want %d, got %d", 4, val)
	}
}
