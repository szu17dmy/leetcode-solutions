package binary_tree_zigzag_level_order_traversal

import (
	"reflect"
	"testing"
)

func Test_zigzagLevelOrder(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name    string
		args    args
		wantAns [][]int
	}{
		{
			name: "Case 1",
			args: args{
				root: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 9,
					},
					Right: &TreeNode{
						Val: 20,
						Left: &TreeNode{
							Val: 15,
						},
						Right: &TreeNode{
							Val: 7,
						},
					},
				},
			},
			wantAns: [][]int{
				{3},
				{20, 9},
				{15, 7},
			},
		},
		{
			name: "Case 2",
			args: args{
				root: &TreeNode{
					Val: 1,
				},
			},
			wantAns: [][]int{
				{1},
			},
		},
		{
			name: "Case 3",
			args: args{
				root: nil,
			},
			wantAns: nil,
		},
		{
			name: "Case 19/33",
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 4,
						},
					},
					Right: &TreeNode{
						Val: 3,
						Right: &TreeNode{
							Val: 5,
						},
					},
				},
			},
			wantAns: [][]int{
				{1},
				{3, 2},
				{4, 5},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := zigzagLevelOrder(tt.args.root); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("zigzagLevelOrder() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
