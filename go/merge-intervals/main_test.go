package merge_intervals

import (
	"reflect"
	"testing"
)

func Test_merge(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Case 1",
			args: args{
				intervals: [][]int{
					{1, 3},
					{2, 6},
					{8, 10},
					{15, 18},
				},
			},
			want: [][]int{
				{1, 6},
				{8, 10},
				{15, 18},
			},
		},
		{
			name: "Case 2",
			args: args{
				intervals: [][]int{
					{1, 4},
					{4, 5},
				},
			},
			want: [][]int{
				{1, 5},
			},
		},
		{
			name: "Case 13/170",
			args: args{
				intervals: [][]int{
					{1, 4},
					{2, 3},
				},
			},
			want: [][]int{
				{1, 4},
			},
		},
		{
			name: "Case 13/170",
			args: args{
				intervals: [][]int{
					{1, 4},
					{0, 0},
				},
			},
			want: [][]int{
				{0, 0},
				{1, 4},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := merge(tt.args.intervals); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("merge() = %v, want %v", got, tt.want)
			}
		})
	}
}
