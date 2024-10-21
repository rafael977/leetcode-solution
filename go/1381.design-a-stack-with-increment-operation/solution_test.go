package designastackwithincrementoperation

import (
	"reflect"
	"testing"
)

func TestConstructor(t *testing.T) {
	type args struct {
		stack CustomStack
		ops   func(CustomStack) []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "tc1",
			args: args{
				stack: Constructor(3),
				ops: func(s CustomStack) (r []int) {
					s.Push(1)
					s.Push(2)
					r = append(r, s.Pop())
					s.Push(2)
					s.Push(3)
					s.Push(4)
					s.Increment(5, 100)
					s.Increment(2, 100)
					r = append(r, s.Pop())
					r = append(r, s.Pop())
					r = append(r, s.Pop())
					r = append(r, s.Pop())
					return
				},
			},
			want: []int{2, 103, 202, 201, -1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.ops(tt.args.stack); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
