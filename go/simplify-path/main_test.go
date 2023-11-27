package simplify_path

import "testing"

func Test_simplifyPath(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Case 1",
			args: args{
				path: "/home/",
			},
			want: "/home",
		},
		{
			name: "Case 2",
			args: args{
				path: "/../",
			},
			want: "/",
		},
		{
			name: "Case 3",
			args: args{
				path: "/home//foo/",
			},
			want: "/home/foo",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := simplifyPath(tt.args.path); got != tt.want {
				t.Errorf("simplifyPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
