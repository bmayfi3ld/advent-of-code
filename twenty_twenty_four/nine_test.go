package twentytwentyfour

import (
	"reflect"
	"testing"
)

func Test_expandFormat(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"example",
			args{"2333133121414131402"},
			[]int{0, 0, -1, -1, -1, 1, 1, 1, -1, -1, -1, 2, -1, -1, -1, 3, 3, 3, -1, 4, 4, -1, 5, 5, 5, 5, -1, 6, 6, 6, 6, -1, 7, 7, 7, -1, 8, 8, 8, 8, 9, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := expandDiskFormat(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("expandFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}
