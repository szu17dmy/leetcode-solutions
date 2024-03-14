package _3sum

import (
	"reflect"
	"testing"
)

func Test_threeSum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name  string
		args  args
		wantR [][]int
	}{
		{
			name: "Case 1",
			args: args{
				nums: []int{-1, 0, 1, 2, -1, -4},
			},
			wantR: [][]int{
				{-1, -1, 2},
				{-1, 0, 1},
			},
		},
		{
			name: "Case 2",
			args: args{
				nums: []int{0, 1, 1},
			},
			wantR: nil,
		},
		{
			name: "Case 3",
			args: args{
				nums: []int{0, 0, 0},
			},
			wantR: [][]int{
				{0, 0, 0},
			},
		},
		{
			name: "Case 132/313",
			args: args{
				nums: []int{-2, 0, 0, 2, 2},
			},
			wantR: [][]int{
				{-2, 0, 2},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotR := threeSum(tt.args.nums); !reflect.DeepEqual(gotR, tt.wantR) {
				t.Errorf("threeSum() = %v, want %v", gotR, tt.wantR)
			}
		})
	}
}
