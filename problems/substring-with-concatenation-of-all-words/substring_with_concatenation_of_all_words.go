package substring_with_concatenation_of_all_words

/*
Because all words are of same length, we can have a function.
f(i)=j such that, s[i,...,i+m] = words[j].
Then we check whether i is an answer by checking f(i) && f(i+m) && ... && f(i+ (n-1)*m)

How to handle duplicates in words?
Record a wordCntMap. where wordCntMap[i]=k means words[i] repeated k times.
Check whether i is an answer by checking whether the cnt map from f(i), f(i+m), ... is the same as the wordCntMap.
*/
func findSubstring(s string, words []string) []int {
	l := len(s)
	n := len(words)
	m := len(words[0])
	f := make([]int, l-m+1)
	wordMap := make(map[string]int)
	wordCntMap := make(map[int]int)
	for i, word := range words {
		if v, ok := wordMap[word]; !ok {
			wordMap[word] = i
			wordCntMap[i] = 1
		} else {
			wordCntMap[v]++
		}
	}

	for i := 0; i <= l-m; i++ {
		v, ok := wordMap[s[i:i+m]]
		if !ok {
			f[i] = -1
			continue
		}
		f[i] = v
	}

	results := []int{}
	for i := 0; i <= l-n*m; i++ {
		cntMap := make(map[int]int)
		j := 0
		for ; j < n; j++ {
			index := f[i+j*m]
			if index == -1 {
				break
			}
			cntMap[index]++
		}
		if j == n && isMapEqual(wordCntMap, cntMap) {
			results = append(results, i)
		}
	}
	return results
}

func isMapEqual(a, b map[int]int) bool {
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if b[k] != v {
			return false
		}
	}
	return true
}
