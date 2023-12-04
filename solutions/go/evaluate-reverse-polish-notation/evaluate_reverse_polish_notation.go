package evaluate_reverse_polish_notation

import "strconv"

// https://leetcode.com/problems/evaluate-reverse-polish-notation

func evalRPN(tokens []string) int {
	var numbers []int
	for _, s := range tokens {
		l := len(numbers)
		if s == "*" || s == "+" || s == "-" || s == "/" {
			a, b := numbers[l-2], numbers[l-1]
			numbers = numbers[:l-1]
			switch s {
			case "*":
				numbers[l-2] = a * b
			case "+":
				numbers[l-2] = a + b
			case "-":
				numbers[l-2] = a - b
			case "/":
				numbers[l-2] = a / b
			}
			continue
		}
		a, _ := strconv.Atoi(s)
		numbers = append(numbers, a)
	}
	return numbers[0]
}
