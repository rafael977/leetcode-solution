package findthelengthofthelongestcommonprefix

import "testing"

func Test_longestCommonPrefix(t *testing.T) {
	type args struct {
		arr1 []int
		arr2 []int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "tc1",
			args: args{
				arr1: []int{1, 10, 100},
				arr2: []int{1000},
			},
			wantAns: 3,
		},
		{
			name: "tc2",
			args: args{
				arr1: []int{1, 2, 3},
				arr2: []int{4, 4, 4},
			},
			wantAns: 0,
		},
		{
			name: "tc3",
			args: args{
				arr1: []int{13, 27, 45},
				arr2: []int{21, 27, 48},
			},
			wantAns: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := longestCommonPrefix(tt.args.arr1, tt.args.arr2); gotAns != tt.wantAns {
				t.Errorf("longestCommonPrefix() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
