package longest_valid_parentheses

// https://leetcode.com/problems/longest-valid-parentheses

/*
Define f[i]:=j where the longest valid parentheses ended at s[i] starts at s[i-j+1]. The answer will be max_i {f[i]}.

There are two cases of composing a longer valid substring from smaller substrings.
1. ([...i-1]). In this case, f[i] = f[i-1] + 2
2. [...i-k][i-k+1...i]. In this case, f[i] = f[i-k] + k. The k is the *shortest* non-empty valid parentheses ends at s[i.]
*/
func longestValidParentheses(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}

	max := 0
	f := make([]int, n)
	f[0] = 0

	matchCheck := func(i, j int) bool {
		return i >= 0 && i < n && j >= 0 && j < n && i < j && s[i] == '(' && s[j] == ')'
	}

	shortestLength := func(i int) int {
		cnt := 0
		j := i
		for ; j >= 0; j-- {
			if s[j] == ')' {
				cnt++
			} else {
				cnt--
			}
			if cnt < 0 {
				break
			}
			if cnt == 0 {
				break
			}
		}
		if cnt != 0 {
			return 0
		}
		return i - j + 1
	}

	for i := 1; i < n; i++ {
		// case 1.
		if f[i-1] > 0 && matchCheck(i-f[i-1]-1, i) {
			f[i] = f[i-1] + 2
		}
		// case 2.
		k := shortestLength(i)
		if i-k+1 == 0 {
			f[i] = k
		} else if k > 0 && i-k >= 0 && f[i-k]+k > f[i] {
			f[i] = f[i-k] + k
		}

		if f[i] > max {
			max = f[i]
		}
	}
	return max
}
