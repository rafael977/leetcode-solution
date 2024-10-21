package isgraphbipartite

import "testing"

func Test_isBipartite(t *testing.T) {
	type args struct {
		graph [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "tc1",
			args: args{
				graph: [][]int{{1, 2, 3}, {0, 2}, {0, 1, 3}, {0, 2}},
			},
			want: false,
		},
		{
			name: "tc2",
			args: args{
				graph: [][]int{{1, 3}, {0, 2}, {1, 3}, {0, 2}},
			},
			want: true,
		},
		{
			name: "tc3",
			args: args{
				graph: [][]int{{2, 4}, {2, 3, 4}, {0, 1}, {1}, {0, 1}, {7}, {9}, {5}, {}, {6}, {12, 14}, {}, {10}, {}, {10}, {19}, {18}, {}, {16}, {15}, {23}, {23}, {}, {20, 21}, {}, {}, {27}, {26}, {}, {}, {34}, {33, 34}, {}, {31}, {30, 31}, {38, 39}, {37, 38, 39}, {36}, {35, 36}, {35, 36}, {43}, {}, {}, {40}, {}, {49}, {47, 48, 49}, {46, 48, 49}, {46, 47, 49}, {45, 46, 47, 48}},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isBipartite(tt.args.graph); got != tt.want {
				t.Errorf("isBipartite() = %v, want %v", got, tt.want)
			}
		})
	}
}
