package main

import (
	"reflect"
	"testing"
)

func Test_cloneGraph(t *testing.T) {
	var buildGraph func([][]int) *Node = func(arr [][]int) *Node {
		nodes := make([]*Node, len(arr))
		for i := range arr {
			nodes[i] = &Node{Val: i + 1, Neighbors: []*Node{}}
		}
		for i, ns := range arr {
			node := nodes[i]
			for _, n := range ns {
				node.Neighbors = append(node.Neighbors, nodes[n-1])
			}
		}
		return nodes[0]
	}

	type args struct {
		node *Node
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		{
			name: "tc1",
			args: args{
				node: buildGraph([][]int{{2, 4}, {1, 3}, {2, 4}, {1, 3}}),
			},
			want: buildGraph([][]int{{2, 4}, {1, 3}, {2, 4}, {1, 3}}),
		},
		{
			name: "tc2",
			args: args{
				node: buildGraph([][]int{{}}),
			},
			want: buildGraph([][]int{{}}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cloneGraph(tt.args.node); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("cloneGraph() = %v, want %v", got, tt.want)
			}
		})
	}
}
