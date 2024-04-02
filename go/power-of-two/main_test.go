package power_of_two

import "testing"

func Test_isPowerOfTwo(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Case 1",
			args: args{
				n: 1,
			},
			want: true,
		},
		{
			name: "Case 2",
			args: args{
				n: 16,
			},
			want: true,
		},
		{
			name: "Case 3",
			args: args{
				n: 3,
			},
			want: false,
		},
		{
			name: "Case 1109/1110",
			args: args{
				n: 0,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPowerOfTwo(tt.args.n); got != tt.want {
				t.Errorf("isPowerOfTwo() = %v, want %v", got, tt.want)
			}
		})
	}
}
