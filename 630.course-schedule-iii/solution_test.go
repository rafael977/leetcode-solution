package coursescheduleiii

import "testing"

func Test_scheduleCourse(t *testing.T) {
	type args struct {
		courses [][]int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "tc1",
			args: args{
				courses: [][]int{{100, 200}, {200, 1300}, {1000, 1250}, {2000, 3200}},
			},
			wantAns: 3,
		},
		{
			name: "tc2",
			args: args{
				courses: [][]int{{1, 2}},
			},
			wantAns: 1,
		},
		{
			name: "tc3",
			args: args{
				courses: [][]int{{3, 2}, {4, 3}},
			},
			wantAns: 0,
		},
		{
			name: "tc4",
			args: args{
				courses: [][]int{{100, 200}, {200, 1300}, {1150, 1250}, {2000, 3200}},
			},
			wantAns: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := scheduleCourse(tt.args.courses); gotAns != tt.wantAns {
				t.Errorf("scheduleCourse() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
