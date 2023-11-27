package evaluate_reverse_polish_notation

import "strconv"

func evalRPN(tokens []string) int {
	var stack []int
	for _, token := range tokens {
		switch token {
		case "+", "-", "*", "/":
			stack[len(stack)-2] = operate(stack[len(stack)-2], stack[len(stack)-1], token)
			stack = stack[:len(stack)-1]
		default:
			i, _ := strconv.Atoi(token)
			stack = append(stack, i)
		}
	}
	return stack[len(stack)-1]
}

func operate(a, b int, operator string) int {
	switch operator {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	}
	return 0
}
