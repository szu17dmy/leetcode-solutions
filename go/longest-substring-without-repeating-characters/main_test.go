package longest_substring_without_repeating_characters

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "Case 1",
			args: args{
				s: "abcabcbb",
			},
			wantAns: 3,
		},
		{
			name: "Case 2",
			args: args{
				s: "bbbbb",
			},
			wantAns: 1,
		},
		{
			name: "Case 3",
			args: args{
				s: "pwwkew",
			},
			wantAns: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := lengthOfLongestSubstring(tt.args.s); gotAns != tt.wantAns {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
