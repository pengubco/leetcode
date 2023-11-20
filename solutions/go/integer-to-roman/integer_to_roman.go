package integertoroman

func intToRoman(num int) string {
	var result []byte
	lookup := map[int]string{
		1:    "I",
		5:    "V",
		10:   "X",
		50:   "L",
		100:  "C",
		500:  "D",
		1000: "M",
		4:    "IV",
		9:    "IX",
		40:   "XL",
		90:   "XC",
		400:  "CD",
		900:  "CM",
	}
	nums := []int{
		1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1,
	}
	for num > 0 {
		for _, x := range nums {
			if num >= x {
				num -= x
				result = append(result, []byte(lookup[x])...)
				break
			}
		}
	}
	return string(result)
}
