package best_time_to_buy_and_sell_stock_ii

import "testing"

func Test_maxProfit(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name  string
		args  args
		wantP int
	}{
		{
			name: "Case 1",
			args: args{
				prices: []int{7, 1, 5, 3, 6, 4},
			},
			wantP: 7,
		},
		{
			name: "Case 2",
			args: args{
				prices: []int{1, 2, 3, 4, 5},
			},
			wantP: 4,
		},
		{
			name: "Case 3",
			args: args{
				prices: []int{7, 6, 4, 3, 1},
			},
			wantP: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotP := maxProfit(tt.args.prices); gotP != tt.wantP {
				t.Errorf("maxProfit() = %v, want %v", gotP, tt.wantP)
			}
		})
	}
}
