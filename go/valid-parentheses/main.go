package valid_parentheses

var t = map[rune]rune{
	')': '(',
	']': '[',
	'}': '{',
}

func isValid(s string) bool {
	var stack []rune
	for _, el := range s {
		if el == '(' || el == '[' || el == '{' {
			stack = append(stack, el)
		} else {
			if len(stack) == 0 {
				return false
			}
			if t[el] != stack[len(stack)-1] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	if len(stack) == 0 {
		return true
	}
	return false
}
