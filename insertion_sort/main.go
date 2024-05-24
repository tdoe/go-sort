package main

import "fmt"

func main() {

	digits := []int{5, 2, 4, 6, 1, 3}

	for i := 1; i < len(digits); i++ {
		tmp := digits[i]

		for j := i - 1; j >= 0; j-- {
			if tmp < digits[j] {
				digits[j+1] = digits[j]
				digits[j] = tmp
			}
		}
	}

	fmt.Println(digits)
}
