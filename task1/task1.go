package task1

import "strconv"

func Editor(num int) string {
	if num%10 == 1 && num%100 != 11 {
		return strconv.Itoa(num) + " компьютер"
	} else if (2 <= num%10 && num%10 <= 4) && (num%100 < 10 || num%100 >= 20) {
		return strconv.Itoa(num) + " компьютера"
	} else {
		return strconv.Itoa(num) + " компьютеров"
	}
}
