package designhashset

import "testing"

func TestMyHashSet(t *testing.T) {
	hset := Constructor()
	hset.Add(1)
	hset.Add(2)
	if !hset.Contains(1) {
		t.Errorf("1 not in set")
	}
	if hset.Contains(3) {
		t.Errorf("3 in set")
	}
	hset.Add(2)
	if !hset.Contains(2) {
		t.Errorf("2 not in set")
	}
	hset.Remove(2)
	if hset.Contains(2) {
		t.Errorf("2 in set")
	}
}
