package summary_ranges

import (
	"reflect"
	"testing"
)

func Test_summaryRanges(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name  string
		args  args
		wantR []string
	}{
		{
			name: "Case 1",
			args: args{
				nums: []int{0, 1, 2, 4, 5, 7},
			},
			wantR: []string{"0->2", "4->5", "7"},
		},
		{
			name: "Case 2",
			args: args{
				nums: []int{0, 2, 3, 4, 6, 8, 9},
			},
			wantR: []string{"0", "2->4", "6", "8->9"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotR := summaryRanges(tt.args.nums); !reflect.DeepEqual(gotR, tt.wantR) {
				t.Errorf("summaryRanges() = %v, want %v", gotR, tt.wantR)
			}
		})
	}
}
