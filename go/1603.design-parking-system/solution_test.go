package designparkingsystem

import "testing"

func TestParkingSystem_AddCar(t *testing.T) {
	type args struct {
		sys      ParkingSystem
		carTypes []int
	}
	tests := []struct {
		name string
		args args
		want []bool
	}{
		{
			name: "tc1",
			args: args{
				sys:      Constructor(1, 1, 0),
				carTypes: []int{1, 2, 3, 1},
			},
			want: []bool{true, true, false, false},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i, v := range tt.args.carTypes {
				res := tt.args.sys.AddCar(v)
				if res != tt.want[i] {
					t.Errorf("At %d, AddCar() = %v, expected = %v", i, res, tt.want[i])
				}
			}
		})
	}
}
