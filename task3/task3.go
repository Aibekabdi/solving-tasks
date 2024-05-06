package task3

func primeNumbers(start, end int) []int {
	if end < start {
		return []int{}
	}
	var res []int
	for i := start; i <= end; i++ {
		if isPrime(i) {
			res = append(res, i)
		}
	}
	return res
}

func isPrime(num int) bool {
	if num <= 1 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
