package _4sum

import (
	"reflect"
	"testing"
)

func Test_fourSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name    string
		args    args
		wantAns [][]int
	}{
		{
			name: "Case 1",
			args: args{
				nums:   []int{1, 0, -1, 0, -2, 2},
				target: 0,
			},
			wantAns: [][]int{
				{-2, -1, 1, 2},
				{-2, 0, 0, 2},
				{-1, 0, 0, 1},
			},
		},
		{
			name: "Case 2",
			args: args{
				nums:   []int{2, 2, 2, 2, 2},
				target: 8,
			},
			wantAns: [][]int{
				{2, 2, 2, 2},
			},
		},
		{
			name: "Case 227/294",
			args: args{
				nums:   []int{-2, -1, -1, 1, 1, 2, 2},
				target: 0,
			},
			wantAns: [][]int{
				{-2, -1, 1, 2},
				{-1, -1, 1, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := fourSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("fourSum() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
