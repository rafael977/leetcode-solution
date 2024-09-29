package alloonedatastructure

import (
	"reflect"
	"testing"
)

func TestConstructor(t *testing.T) {
	type args struct {
		allOne AllOne
		steps  func(AllOne) []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "tc1",
			args: args{
				allOne: Constructor(),
				steps: func(a AllOne) (r []string) {
					a.Inc("hello")
					a.Inc("hello")
					r = append(r, a.GetMaxKey())
					r = append(r, a.GetMinKey())
					a.Inc("leet")
					r = append(r, a.GetMaxKey())
					r = append(r, a.GetMinKey())
					return
				},
			},
			want: []string{"hello", "hello", "hello", "leet"},
		},
		{
			name: "tc2",
			args: args{
				allOne: Constructor(),
				steps: func(a AllOne) (r []string) {
					a.Inc("a")
					a.Inc("b")
					a.Inc("b")
					a.Inc("b")
					a.Dec("b")
					a.Dec("b")
					r = append(r, a.GetMaxKey())
					r = append(r, a.GetMinKey())
					return
				},
			},
			want: []string{"a", "a"},
		},
		{
			name: "tc3",
			args: args{
				allOne: Constructor(),
				steps: func(a AllOne) (r []string) {
					a.Inc("a")
					a.Inc("b")
					a.Inc("b")
					a.Inc("c")
					a.Inc("c")
					a.Inc("c")
					a.Dec("b")
					a.Dec("b")
					r = append(r, a.GetMinKey())
					a.Dec("a")
					r = append(r, a.GetMaxKey())
					r = append(r, a.GetMinKey())
					return
				},
			},
			want: []string{"a", "c", "c"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.steps(tt.args.allOne); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got: %v, expect: %v", got, tt.want)
			}
		})
	}
}
