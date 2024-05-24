package main

import "fmt"

func main() {
	digits := []int{5, 2, 4, 6, 1, 3}
	for i, di := range digits {
		for j, dj := range digits {
			if di < dj {
				digits[i], digits[j] = digits[j], digits[i]
			}
		}
	}

	fmt.Println(digits)
}
