package validatebinarytreenodes

import "testing"

func Test_validateBinaryTreeNodes(t *testing.T) {
	type args struct {
		n          int
		leftChild  []int
		rightChild []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "tc1",
			args: args{
				n:          4,
				leftChild:  []int{1, -1, 3, -1},
				rightChild: []int{2, -1, -1, -1},
			},
			want: true,
		},
		{
			name: "tc2",
			args: args{
				n:          4,
				leftChild:  []int{1, -1, 3, -1},
				rightChild: []int{2, 3, -1, -1},
			},
			want: false,
		},
		{
			name: "tc3",
			args: args{
				n:          2,
				leftChild:  []int{1, 0},
				rightChild: []int{-1, -1},
			},
			want: false,
		},
		{
			name: "tc4",
			args: args{
				n:          6,
				leftChild:  []int{1, -1, -1, 4, -1, -1},
				rightChild: []int{2, -1, -1, 5, -1, -1},
			},
			want: false,
		},
		{
			name: "tc5",
			args: args{
				n:          4,
				leftChild:  []int{3, -1, 1, -1},
				rightChild: []int{-1, -1, 0, -1},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validateBinaryTreeNodes(tt.args.n, tt.args.leftChild, tt.args.rightChild); got != tt.want {
				t.Errorf("validateBinaryTreeNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}
