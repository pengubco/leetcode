package romantointeger

func romanToInt(s string) int {
	lookup := map[string]int{
		"I":  1,
		"V":  5,
		"X":  10,
		"L":  50,
		"C":  100,
		"D":  500,
		"M":  1000,
		"IV": 4,
		"IX": 9,
		"XL": 40,
		"XC": 90,
		"CD": 400,
		"CM": 900,
	}

	n := len(s)
	result := 0
	for i := 0; i < n; {
		if i+1 < n && lookup[s[i:i+2]] > 0 {
			result += lookup[s[i:i+2]]
			i += 2
			continue
		}
		result += lookup[s[i:i+1]]
		i++
		continue
	}
	return result
}
