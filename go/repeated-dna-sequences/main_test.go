package repeated_dna_sequences

import (
	"reflect"
	"testing"
)

func Test_findRepeatedDnaSequences(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name       string
		args       args
		wantResult []string
	}{
		{
			name: "Case 1",
			args: args{
				s: "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT",
			},
			wantResult: []string{"AAAAACCCCC", "CCCCCAAAAA"},
		},
		{
			name: "Case 2",
			args: args{
				s: "AAAAAAAAAAAAA",
			},
			wantResult: []string{"AAAAAAAAAA"},
		},
		{
			name: "Case 25",
			args: args{
				s: "AAAAAAAAAAA",
			},
			wantResult: []string{"AAAAAAAAAA"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := findRepeatedDnaSequences(tt.args.s); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("findRepeatedDnaSequences() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
