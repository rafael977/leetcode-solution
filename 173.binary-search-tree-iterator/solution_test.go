package binarysearchtreeiterator

import (
	"testing"

	"github.com/rafael977/leetcode-solution/datastruct"
)

func TestBSTIterator(t *testing.T) {
	it := Constructor(datastruct.BuildTree("7, 3, 15, null, null, 9, 20"))
	assertValue(t, it.Next(), 3)
	assertValue(t, it.Next(), 7)
	assertValue(t, it.HasNext(), true)
	assertValue(t, it.Next(), 9)
	assertValue(t, it.HasNext(), true)
	assertValue(t, it.Next(), 15)
	assertValue(t, it.HasNext(), true)
	assertValue(t, it.Next(), 20)
	assertValue(t, it.HasNext(), false)
}

func assertValue[T int | bool](t *testing.T, got, want T) {
	if got != want {
		t.Errorf("got: %v, want: %v", got, want)
	}
}
