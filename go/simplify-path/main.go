package simplify_path

import "strings"

func simplifyPath(path string) string {
	list := strings.Split(path, "/")
	var stack []string
	for _, el := range list {
		switch el {
		case ".", "":
		case "..":
			if len(stack) != 0 {
				stack = stack[:len(stack)-1]
			}
		default:
			stack = append(stack, el)
		}
	}
	return "/" + strings.Join(stack, "/")
}
