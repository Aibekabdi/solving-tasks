package task4

import "fmt"

func multTable(a int) {
	for i := 1; i <= a; i++ {
		if i == 1 {
			fmt.Print("  ")
		}
		if i == a {
			fmt.Printf("%d\n", i)
		} else {
			fmt.Printf("%d ", i)
		}
	}
	for i := 1; i <= a; i++ {
		fmt.Printf("%d ", i)
		for j := 1; j <= a; j++ {
			if j == a {
				fmt.Printf("%d\n", i*j)
			} else {
				fmt.Printf("%d ", i*j)
			}
		}
	}
}
