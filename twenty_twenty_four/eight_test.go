package twentytwentyfour

import (
	"fmt"
	"testing"
)

func Test_calcGridMin(t *testing.T) {
	tests := []struct {
		cur  int
		want int
	}{
		{1, 1},
		{2, 1},
		{3, 2},
		{8, 4},
		{9, 5},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.cur), func(t *testing.T) {
			if got := calcGridMin(tt.cur); got != tt.want {
				t.Errorf("calcGridMin(%v) = %v, want %v", tt.cur, got, tt.want)
			}
		})
	}
}

func Test_calcGridMax(t *testing.T) {
	tests := []struct {
		cur  int
		len  int
		want int
	}{
		{2, 9, 5},
		{8, 12, 9},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.cur), func(t *testing.T) {
			if got := calcGridMax(tt.cur, tt.len); got != tt.want {
				t.Errorf("calcGridMin(%v, %v) = %v, want %v", tt.cur, tt.len, got, tt.want)
			}
		})
	}
}
