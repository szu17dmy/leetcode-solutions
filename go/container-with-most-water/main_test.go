package container_with_most_water

import "testing"

func Test_maxArea(t *testing.T) {
	type args struct {
		height []int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "Case 1",
			args: args{
				height: []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			},
			wantAns: 49,
		},
		{
			name: "Case 2",
			args: args{
				height: []int{1, 1},
			},
			wantAns: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := maxArea(tt.args.height); gotAns != tt.wantAns {
				t.Errorf("maxArea() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
