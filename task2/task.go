package task2

import "sort"

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func commonDivisors(arr []int) []int {
	g := arr[0]
	for i := 1; i < len(arr); i++ {
		g = gcd(arr[i], g)
	}
	res := []int{}
	for i := 1; i*i <= g; i++ {
		if g%i == 0 {
			if i != 1 {
				res = append(res, i)
			}
			if g/i != i {
				res = append(res, g/i)
			}
		}
	}
	sort.Ints(res)
	return res
}
