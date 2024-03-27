package exclusive_time_of_functions

import (
	"strconv"
	"strings"
)

func exclusiveTime(n int, logs []string) []int {
	var stack [][]int
	ans := make([]int, n)
	for _, log := range logs {
		pid, act, t := parse(log)
		// Action == "start"
		if act[0] == 's' {
			if len(stack) > 0 {
				last := stack[len(stack)-1]
				ans[last[0]] += t - last[1]
				last[1] = t
			}
			stack = append(stack, []int{pid, t})
			continue
		}
		// Action == "end"
		last := stack[len(stack)-1]
		ans[pid] += t - last[1] + 1
		stack = stack[:len(stack)-1]
		if len(stack) > 0 {
			stack[len(stack)-1][1] = t + 1
		}
	}
	return ans
}

func parse(log string) (int, string, int) {
	s := strings.Split(log, ":")
	if len(s) != 3 {
		panic(s)
	}
	p, _ := strconv.Atoi(s[0])
	t, _ := strconv.Atoi(s[2])
	return p, s[1], t
}
