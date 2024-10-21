package maximumunitsonatruck

import "testing"

func Test_maximumUnits(t *testing.T) {
	type args struct {
		boxTypes  [][]int
		truckSize int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "tc1",
			args: args{
				boxTypes:  [][]int{{1, 3}, {2, 2}, {3, 1}},
				truckSize: 4,
			},
			wantAns: 8,
		},
		{
			name: "tc2",
			args: args{
				boxTypes:  [][]int{{5, 10}, {2, 5}, {4, 7}, {3, 9}},
				truckSize: 10,
			},
			wantAns: 91,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := maximumUnits(tt.args.boxTypes, tt.args.truckSize); gotAns != tt.wantAns {
				t.Errorf("maximumUnits() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
