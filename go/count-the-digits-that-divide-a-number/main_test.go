package count_the_digits_that_divide_a_number

import "testing"

func Test_countDigits(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name  string
		args  args
		wantR int
	}{
		{
			name: "Case 1",
			args: args{
				num: 7,
			},
			wantR: 1,
		},
		{
			name: "Case 2",
			args: args{
				num: 121,
			},
			wantR: 2,
		},
		{
			name: "Case 3",
			args: args{
				num: 1248,
			},
			wantR: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotR := countDigits(tt.args.num); gotR != tt.wantR {
				t.Errorf("countDigits() = %v, want %v", gotR, tt.wantR)
			}
		})
	}
}
