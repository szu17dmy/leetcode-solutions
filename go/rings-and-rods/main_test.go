package rings_and_rods

import "testing"

func Test_countPoints(t *testing.T) {
	type args struct {
		rings string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Case 1",
			args: args{
				rings: "B0B6G0R6R0R6G9",
			},
			want: 1,
		},
		{
			name: "Case 2",
			args: args{
				rings: "B0R0G0R9R0B0G0",
			},
			want: 1,
		},
		{
			name: "Case 3",
			args: args{
				rings: "G4",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countPoints(tt.args.rings); got != tt.want {
				t.Errorf("countPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}
