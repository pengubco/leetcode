package longestrepeatingcharacterreplacement

// https://leetcode.com/problems/longest-repeating-character-replacement/

func characterReplacement(s string, k int) int {
	n := len(s)
	sumMap := make(map[byte]int)
	maxValueOfMap := func() int {
		result := 0
		for _, v := range sumMap {
			if v > result {
				result = v
			}
		}
		return result
	}

	result := 0
	lastJ := -1
	for i, j := 0, 0; i < n && j < n; {
		for ; j < n; j++ {
			if j > lastJ {
				sumMap[s[j]]++
				lastJ = j
			}
			if (j-i+1)-maxValueOfMap() > k {
				break
			}
			cnt := j - i + 1
			if cnt > result {
				result = cnt
			}
		}
		sumMap[s[i]]--
		i++
	}
	return result
}
