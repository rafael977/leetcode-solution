package mycalendarii

import "testing"

func TestMyCalendarTwo_Book(t *testing.T) {
	c := Constructor()
	type args struct {
		start int
		end   int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "tc1",
			args: args{
				start: 10,
				end:   20,
			},
			want: true,
		},
		{
			name: "tc2",
			args: args{
				start: 50,
				end:   60,
			},
			want: true,
		},
		{
			name: "tc3",
			args: args{
				start: 10,
				end:   40,
			},
			want: true,
		},
		{
			name: "tc4",
			args: args{
				start: 5,
				end:   15,
			},
			want: false,
		},
		{
			name: "tc5",
			args: args{
				start: 5,
				end:   10,
			},
			want: true,
		},
		{
			name: "tc5",
			args: args{
				start: 25,
				end:   55,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := c.Book(tt.args.start, tt.args.end); got != tt.want {
				t.Errorf("MyCalendarTwo.Book() = %v, want %v", got, tt.want)
			}
		})
	}
}
