package longestrepeatingcharacterreplacement

/*
If we know the repeating character is 'A', then this question is the same as
https://leetcode.com/problems/max-consecutive-ones-iii/description/
*/
func characterReplacement(s string, k int) int {
	n := len(s)
	// 2-d map: sumMap[c][i]: count of c in s[:i+1]
	sumMap := make(map[byte]map[int]int)
	for i := 0; i < n; i++ {
		c := s[i]
		sum, ok := sumMap[c]
		if !ok {
			sum = make(map[int]int)
			sumMap[c] = sum
		}
		sum[i] = sum[i-1] + 1
	}
}
