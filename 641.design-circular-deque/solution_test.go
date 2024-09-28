package designcirculardeque

import (
	"testing"
)

func Test_MyCircularDeque(t *testing.T) {
	q := Constructor(3)

	tests := []struct {
		name   string
		assert func() bool
	}{
		{
			name: "tc1",
			assert: func() bool {
				return q.InsertLast(1) == true
			},
		},
		{
			name: "tc2",
			assert: func() bool {
				return q.InsertLast(2) == true
			},
		},
		{
			name: "tc3",
			assert: func() bool {
				return q.InsertFront(3) == true
			},
		},
		{
			name: "tc4",
			assert: func() bool {
				return q.InsertFront(4) == false
			},
		},
		{
			name: "tc5",
			assert: func() bool {
				return q.GetRear() == 2
			},
		},
		{
			name: "tc6",
			assert: func() bool {
				return q.IsFull() == true
			},
		},
		{
			name: "tc7",
			assert: func() bool {
				return q.DeleteLast() == true
			},
		},
		{
			name: "tc8",
			assert: func() bool {
				return q.InsertFront(4) == true
			},
		},
		{
			name: "tc9",
			assert: func() bool {
				return q.GetFront() == 4
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.assert() {
				t.Errorf("wrong result")
			}
		})
	}
}
