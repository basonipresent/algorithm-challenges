package main

import "testing"

func TestReverse(t *testing.T) {
	tests := []struct {
		name string
		x    int
		want int
	}{
		{
			name: "positive number",
			x:    123,
			want: 321,
		},
		{
			name: "negative number",
			x:    -123,
			want: -321,
		},
		{
			name: "trailing zero dropped",
			x:    120,
			want: 21,
		},
		{
			name: "zero",
			x:    0,
			want: 0,
		},
		{
			name: "single digit",
			x:    7,
			want: 7,
		},
		{
			name: "single negative digit",
			x:    -7,
			want: -7,
		},
		{
			name: "32-bit overflow on positive reversal",
			x:    1534236469,
			want: 0,
		},
		{
			name: "32-bit overflow on negative reversal",
			x:    -2147483648,
			want: 0,
		},
		{
			name: "reversal fits exactly within int32 max",
			x:    1463847412,
			want: 2147483641,
		},
		{
			name: "reversal fits exactly within int32 min",
			x:    -2147483648,
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverse(tt.x)
			if got != tt.want {
				t.Errorf("reverse(%d) = %d, want %d", tt.x, got, tt.want)
			}
		})
	}
}
