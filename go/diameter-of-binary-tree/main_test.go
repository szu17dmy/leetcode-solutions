package diameter_of_binary_tree

import "testing"

func Test_diameterOfBinaryTree(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "Case 1",
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 4,
						},
						Right: &TreeNode{
							Val: 5,
						},
					},
					Right: &TreeNode{
						Val: 3,
					},
				},
			},
			wantAns: 3,
		},
		{
			name: "Case 2",
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 2,
					},
				},
			},
			wantAns: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := diameterOfBinaryTree(tt.args.root); gotAns != tt.wantAns {
				t.Errorf("diameterOfBinaryTree() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
