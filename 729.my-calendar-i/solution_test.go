package mycalendari

import "testing"

func TestMyCalendar_Book(t *testing.T) {
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
			name: "call1",
			args: args{
				start: 10,
				end:   20,
			},
			want: true,
		},
		{
			name: "call2",
			args: args{
				start: 15,
				end:   25,
			},
			want: false,
		},
		{
			name: "call3",
			args: args{
				start: 20,
				end:   30,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := c.Book(tt.args.start, tt.args.end); got != tt.want {
				t.Errorf("MyCalendar.Book() = %v, want %v", got, tt.want)
			}
		})
	}
}
