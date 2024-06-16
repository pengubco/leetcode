package basic_calculator

// https://leetcode.com/problems/basic-calculator/description/

/*
1. Without parenthesis, it is simple because there are only two binary operators '+', '-' and one unary operator '-'.
2. So how do we handle parenthesis, it is a problem of smaller size. So we just use recursive.
*/
func calculate(s string) int {
	var operands []int
	operator := ' '
	n := len(s)
	calcTwoNumbers := func() {
		if len(operands) == 2 && operator == '+' {
			operands[0] = operands[0] + operands[1]
			operands = operands[:1]
			operator = ' '
		} else if len(operands) == 2 && operator == '-' {
			operands[0] = operands[0] - operands[1]
			operands = operands[:1]
			operator = ' '
		}
	}
	for i := 0; i < n; {
		if s[i] == ' ' {
			i++
			continue
		}
		if s[i] == '-' && len(operands) == 0 {
			number, endIndex := readNumber(i, s)
			operands = append(operands, number)
			i = endIndex + 1
			continue
		}
		if s[i] == '(' {
			endIndex := indexOfClosingParenthesis(i, s)
			number := calculate(s[i+1 : endIndex])
			operands = append(operands, number)
			calcTwoNumbers()
			i = endIndex + 1
			continue
		}
		if s[i] == '+' || s[i] == '-' {
			operator = int32(s[i])
			i++
			continue
		}
		if s[i] >= '0' && s[i] <= '9' {
			number, endIndex := readNumber(i, s)
			operands = append(operands, number)
			calcTwoNumbers()
			i = endIndex + 1
			continue
		}
	}

	return operands[0]
}

// read as many characters as needed for a number. return a pair (number, index of the last index of the char)
func readNumber(begin int, s string) (int, int) {
	number := 0
	n := len(s)
	i := begin
	for ; i < n; i++ {
		if s[i] == ' ' {
			continue
		}
		if s[i] >= '0' && s[i] <= '9' {
			number = number*10 + int(s[i]-'0')
			continue
		}
		i--
		break
	}
	return number, i
}

func indexOfClosingParenthesis(begin int, s string) int {
	cnt := 1
	n := len(s)
	for i := begin + 1; i < n; i++ {
		if s[i] == '(' {
			cnt++
		} else if s[i] == ')' {
			cnt--
			if cnt == 0 {
				return i
			}
		}
	}
	return -1
}
