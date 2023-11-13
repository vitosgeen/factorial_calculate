package main

import "testing"

func Test_calculateFactorial(t *testing.T) {
	type args struct {
		number int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "factorial of 0",
			args: args{
				number: 0,
			},
			want: 1,
		},
		{
			name: "factorial of 1",
			args: args{
				number: 1,
			},
			want: 1,
		},
		{
			name: "factorial of 2",
			args: args{
				number: 2,
			},
			want: 2,
		},
		{
			name: "factorial of 3",
			args: args{
				number: 3,
			},
			want: 6,
		},
		{
			name: "factorial of 4",
			args: args{
				number: 4,
			},
			want: 24,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calculateFactorial(tt.args.number)
			if got != tt.want {
				t.Errorf("calculateFactorial() = %v, want %v", got, tt.want)
			}
		})
	}
}
