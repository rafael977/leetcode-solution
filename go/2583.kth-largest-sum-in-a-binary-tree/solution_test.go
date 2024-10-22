// Created by Rafael Shen at 2024/10/22 21:03
// leetgo: 1.4.9
// https://leetcode.com/problems/kth-largest-sum-in-a-binary-tree/

package main

import (
	"testing"

	. "github.com/rafael977/leetcode-solution/datastruct"
)

func Test_kthLargestLevelSum(t *testing.T) {
	type args struct {
		root *TreeNode
		k    int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int64
	}{
		{
			name: "tc1",
			args: args{
				root: BuildTree("5,8,9,2,1,3,7,4,6"),
				k:    2,
			},
			wantAns: 13,
		},
		{
			name: "tc2",
			args: args{
				root: BuildTree("1,2,null,3"),
				k:    1,
			},
			wantAns: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := kthLargestLevelSum(tt.args.root, tt.args.k); gotAns != tt.wantAns {
				t.Errorf("kthLargestLevelSum() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
