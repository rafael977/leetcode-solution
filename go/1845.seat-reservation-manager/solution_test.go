package seatreservationmanager

import (
	"testing"
)

func TestSeatManager(t *testing.T) {
	sm := Constructor(5)
	if out := sm.Reserve(); out != 1 {
		t.Errorf("expect: 1, got: %d", out)
	}
	if out := sm.Reserve(); out != 2 {
		t.Errorf("expect: 2, got: %d", out)
	}
	sm.Unreserve(2)
	if out := sm.Reserve(); out != 2 {
		t.Errorf("expect: 2, got: %d", out)
	}
	if out := sm.Reserve(); out != 3 {
		t.Errorf("expect: 3, got: %d", out)
	}
	if out := sm.Reserve(); out != 4 {
		t.Errorf("expect: 4, got: %d", out)
	}
	if out := sm.Reserve(); out != 5 {
		t.Errorf("expect: 5, got: %d", out)
	}
}
