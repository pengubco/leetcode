package minimum_window_substring

// https://leetcode.com/problems/minimum-window-substring

/*
Define f(i,j) := whether s[i...j] contains all characters from t.
Then f(i,j) can be transitioned from f(i-1, j) and f(i, j+1).
The approach above is O(n^2) because of the two dimension states. Can we do better?

1. What states can we throw away?
1.1. If we know f(i,j) is the smallest substring that starts with s[i]. Then f(i,k), k>j is not needed.
2.2. If we know f(i,j) does not contain all characters. Then f(i,k), k<j is not needed.
*/
func minWindow(s string, t string) string {
	// early exit
	n, m := len(s), len(t)
	if n < m {
		return ""
	}

	targetCnt := make(map[int8]int)
	for _, c := range t {
		targetCnt[int8(c)]++
	}
	var result string
	cnt := make(map[int8]int)
	cnt[int8(s[0])]++
	for i, j := 0, 0; ; {
		if !contains(cnt, targetCnt) {
			j++
			if j >= n {
				break
			}
			cnt[int8(s[j])]++
			continue
		}
		if result == "" || len(result) > j-i+1 {
			result = s[i : j+1]
		}
		cnt[int8(s[i])]--
		i++
		if i > n-m {
			break
		}
	}
	return result
}

func contains(curr, target map[int8]int) bool {
	if len(curr) < len(target) {
		return false
	}
	for k, v := range target {
		if curr[k] < v {
			return false
		}
	}
	return true
}
