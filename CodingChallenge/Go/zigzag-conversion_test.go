package main

import "testing"

func TestConvert(t *testing.T) {
	tests := []struct {
		name    string
		s       string
		numRows int
		want    string
	}{
		{
			name:    "example: 3 rows",
			s:       "PAYPALISHIRING",
			numRows: 3,
			want:    "PAHNAPLSIIGYIR",
		},
		{
			name:    "example: 4 rows",
			s:       "PAYPALISHIRING",
			numRows: 4,
			want:    "PINALSIGYAHRPI",
		},
		{
			name:    "single row",
			s:       "AB",
			numRows: 1,
			want:    "AB",
		},
		{
			name:    "numRows >= len(s)",
			s:       "AB",
			numRows: 5,
			want:    "AB",
		},
		{
			name:    "two rows",
			s:       "ABCDE",
			numRows: 2,
			want:    "ACEBD",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := convert(tt.s, tt.numRows)
			if got != tt.want {
				t.Errorf("convert(%q, %d) = %q, want %q", tt.s, tt.numRows, got, tt.want)
			}
		})
	}
}
