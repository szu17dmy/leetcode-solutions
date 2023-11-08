package find_the_longest_balanced_substring_of_a_binary_string

import "testing"

func Test_findTheLongestBalancedSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name  string
		args  args
		wantR int
	}{
		{
			name: "Case 1",
			args: args{
				s: "01000111",
			},
			wantR: 6,
		},
		{
			name: "Case 2",
			args: args{
				s: "00111",
			},
			wantR: 4,
		},
		{
			name: "Case 3",
			args: args{
				s: "111",
			},
			wantR: 0,
		},
		{
			name: "Case 1231",
			args: args{
				s: "001101",
			},
			wantR: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotR := findTheLongestBalancedSubstring(tt.args.s); gotR != tt.wantR {
				t.Errorf("findTheLongestBalancedSubstring() = %v, want %v", gotR, tt.wantR)
			}
		})
	}
}
