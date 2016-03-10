package romans

import "strings"

var basic map[int]string

func convert(num int) string {
	var digits []string
	base := 1
	for num > 0 {
		digits = append([]string{basic[base*(num%10)]}, digits...)
		base *= 10
		num /= 10
	}
	return strings.Join(digits, "")
}

func init() {
	basic = make(map[int]string)
	basic[1] = "I"
	basic[2] = "II"
	basic[3] = "III"
	basic[4] = "IV"
	basic[5] = "V"
	basic[6] = "VI"
	basic[7] = "VII"
	basic[8] = "VIII"
	basic[9] = "IX"

	basic[10] = "X"
	basic[20] = "XX"
	basic[30] = "XXX"
	basic[40] = "XL"
	basic[50] = "L"
	basic[60] = "LX"
	basic[70] = "LXX"
	basic[80] = "LXXX"
	basic[90] = "XC"

	basic[100] = "C"
	basic[200] = "CC"
	basic[300] = "CCC"
	basic[400] = "CD"
	basic[500] = "D"
	basic[600] = "DC"
	basic[700] = "DCC"
	basic[800] = "DCCC"
	basic[900] = "CM"

	basic[1000] = "M"
	basic[2000] = "MM"
	basic[3000] = "MMM"
}
