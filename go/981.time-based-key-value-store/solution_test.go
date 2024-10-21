package timebasedkeyvaluestore

import "testing"

func Test_TimeMap(t *testing.T) {
	tm := Constructor()
	tm.Set("foo", "bar", 1)
	assert(t, tm.Get("foo", 1), "bar")
	assert(t, tm.Get("foo", 3), "bar")
	tm.Set("foo", "bar2", 4)
	assert(t, tm.Get("foo", 4), "bar2")
	assert(t, tm.Get("foo", 5), "bar2")
	assert(t, tm.Get("foo", 1), "bar")

}

func assert(t *testing.T, got, expect string) {
	if got != expect {
		t.Errorf("got: %v, expect: %v", got, expect)
	}
}
