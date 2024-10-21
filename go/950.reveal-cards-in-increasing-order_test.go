package leetcodesolution

import (
	"reflect"
	"testing"
)

func Test_deckRevealedIncreasing(t *testing.T) {
	type args struct {
		deck []int
	}
	tests := []struct {
		name    string
		args    args
		wantAns []int
	}{
		{
			name: "tc1",
			args: args{
				deck: []int{17, 13, 11, 2, 3, 5, 7},
			},
			wantAns: []int{2, 13, 3, 11, 5, 17, 7},
		},
		{
			name: "tc2",
			args: args{
				deck: []int{1, 1000},
			},
			wantAns: []int{1, 1000},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := deckRevealedIncreasing(tt.args.deck); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("deckRevealedIncreasing() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
