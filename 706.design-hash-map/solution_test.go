package designhashmap

import (
	"testing"
)

func TestMyHashMap(t *testing.T) {
	assertGet := func(got, want int) {
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	}

	hm := Constructor()

	hm.Put(1, 1)
	hm.Put(2, 2)
	assertGet(hm.Get(1), 1)
	assertGet(hm.Get(3), -1)
	hm.Put(2, 1)
	assertGet(hm.Get(2), 1)
	hm.Remove(2)
	assertGet(hm.Get(2), -1)
}
