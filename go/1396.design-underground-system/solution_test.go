package designundergroundsystem

import (
	"math"
	"testing"
)

const EPSILON float64 = 1e-5

func TestUndergroundSystem_1(t *testing.T) {
	sys := Constructor()
	sys.CheckIn(45, "Leyton", 3)
	sys.CheckIn(32, "Paradise", 8)
	sys.CheckIn(27, "Leyton", 10)
	sys.CheckOut(45, "Waterloo", 15)
	sys.CheckOut(27, "Waterloo", 20)
	sys.CheckOut(32, "Cambridge", 22)

	if a := sys.GetAverageTime("Paradise", "Cambridge"); math.Abs(a-14) > EPSILON {
		t.Errorf("Paradise -> Cambridge: got: %f, expected: %f", a, 14.0)
	}
	if a := sys.GetAverageTime("Leyton", "Waterloo"); math.Abs(a-11) > EPSILON {
		t.Errorf("Leyton -> Waterloo: got: %f, expected: %f", a, 11.0)
	}

	sys.CheckIn(10, "Leyton", 24)
	if a := sys.GetAverageTime("Leyton", "Waterloo"); math.Abs(a-11) > EPSILON {
		t.Errorf("Leyton -> Waterloo: got: %f, expected: %f", a, 11.0)
	}
	sys.CheckOut(10, "Waterloo", 38)
	if a := sys.GetAverageTime("Leyton", "Waterloo"); math.Abs(a-12) > EPSILON {
		t.Errorf("Leyton -> Waterloo: got: %f, expected: %f", a, 12.0)
	}
}

func TestUndergroundSystem_2(t *testing.T) {
	sys := Constructor()
	sys.CheckIn(10, "Leyton", 3)
	sys.CheckOut(10, "Paradise", 8)
	if a := sys.GetAverageTime("Leyton", "Paradise"); math.Abs(a-5) > EPSILON {
		t.Errorf("Leyton -> Paradise: got: %f, expected: %f", a, 5.0)
	}
	sys.CheckIn(5, "Leyton", 10)
	sys.CheckOut(5, "Paradise", 16)
	if a := sys.GetAverageTime("Leyton", "Paradise"); math.Abs(a-5.5) > EPSILON {
		t.Errorf("Leyton -> Waterloo: got: %f, expected: %f", a, 5.5)
	}

	sys.CheckIn(2, "Leyton", 21)
	sys.CheckOut(2, "Paradise", 30)
	if a := sys.GetAverageTime("Leyton", "Paradise"); math.Abs(a-6.66667) > EPSILON {
		t.Errorf("Leyton -> Waterloo: got: %f, expected: %f", a, 6.66667)
	}
}
