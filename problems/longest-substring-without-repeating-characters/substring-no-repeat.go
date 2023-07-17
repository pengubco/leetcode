/*
sliding window.
*/
package substring_no_repeat

func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}
	seen := make(map[byte]struct{})
	result := 1
	n := len(s)
	seen[s[0]] = struct{}{}
	for i, j := 0, 1; j < n; {
		if _, ok := seen[s[j]]; !ok {
			seen[s[j]] = struct{}{}
			if j-i+1 > result {
				result = j - i + 1
			}
			j++
			continue
		} else {
			delete(seen, s[i])
			i++
		}
	}
	return result
}
