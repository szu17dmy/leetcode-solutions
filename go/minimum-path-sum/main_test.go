package minimum_path_sum

import "testing"

func Test_minPathSum(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Case 1",
			args: args{
				grid: [][]int{
					{1, 3, 1},
					{1, 5, 1},
					{4, 2, 1},
				},
			},
			want: 7,
		},
		{
			name: "Case 2",
			args: args{
				grid: [][]int{
					{1, 2, 3},
					{4, 5, 6},
				},
			},
			want: 12,
		},
		{
			name: "Case 22/63",
			args: args{
				grid: [][]int{
					{1, 2},
					{5, 6},
					{1, 1},
				},
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minPathSum(tt.args.grid); got != tt.want {
				t.Errorf("minPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
