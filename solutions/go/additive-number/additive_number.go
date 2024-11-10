package additivenumber

import (
	"errors"
	"strconv"
)

// Mind that "0" is valid in the sequence, but "01" is not.

func isAdditiveNumber(s string) bool {
	n := len(s)
	// s[:i], s[i:j] are first two numbers
	for i := 1; i < n; i++ {
		for j := i + 1; j < n; j++ {
			a, b := 0, 0
			var err error
			if a, err = atoi(s[:i]); err != nil {
				continue
			}
			if b, err = atoi(s[i:j]); err != nil {
				continue
			}
			prev1, prev2 := a, b
			k := j
			for k != n {
				c := prev1 + prev2
				next := strconv.Itoa(c)
				if k+len(next) > n || s[k:k+len(next)] != next {
					break
				}
				k += len(next)
				prev1, prev2 = prev2, c
			}
			if k == n {
				return true
			}
		}
	}
	return false
}

func atoi(s string) (int, error) {
	if s == "0" {
		return 0, nil
	}
	if s == "" || s[0] == '0' {
		return 0, errors.New("invalid")
	}
	return strconv.Atoi(s)
}
