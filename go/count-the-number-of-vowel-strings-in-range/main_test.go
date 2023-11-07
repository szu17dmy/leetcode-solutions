package count_the_number_of_vowel_strings_in_range

import "testing"

func Test_vowelStrings(t *testing.T) {
	type args struct {
		words []string
		left  int
		right int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Case 1",
			args: args{
				words: []string{"are", "amy", "u"},
				left:  0,
				right: 2,
			},
			want: 2,
		},
		{
			name: "Case 2",
			args: args{
				words: []string{"hey", "aeo", "mu", "ooo", "artro"},
				left:  1,
				right: 4,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := vowelStrings(tt.args.words, tt.args.left, tt.args.right); got != tt.want {
				t.Errorf("vowelStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
