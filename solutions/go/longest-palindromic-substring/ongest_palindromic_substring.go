package longestpalindromicsubstring

// https://leetcode.com/problems/longest-palindromic-substring/

/*
f[i,j]: is s[i...j] a palindromic substring?
f[i,i] = 1
f[i,j] = f[i+1, j-1] && a[i]=a[j]; j > i
*/
func longestPalindrome(s string) string {
	ans := s[0:1]
	n := len(s)
	f := make([][]bool, n)
	for i := 0; i < n; i++ {
		f[i] = make([]bool, n)
		f[i][i] = true
		if i+1 < n && s[i] == s[i+1] {
			f[i][i+1] = true
			if len(ans) < 2 {
				ans = s[i : i+2]
			}
		}
	}
	for k := 3; k <= n; k++ {
		for i := 0; i+k-1 < n; i++ {
			j := i + k - 1
			if s[i] != s[j] {
				f[i][j] = false
			} else if !f[i+1][j-1] {
				f[i][j] = false
			} else {
				f[i][j] = true
				if len(ans) < k {
					ans = s[i : j+1]
				}
			}
		}
	}
	return ans
}
