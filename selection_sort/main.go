package main

import "fmt"

func main() {
	digits := []int{5, 2, 4, 6, 1, 3}

	for i, d := range digits {
		min, minIndex := d, i
		for j := i; j < len(digits); j++ {
			if min > digits[j] {
				min, minIndex = digits[j], j
			}
		}
		digits[minIndex], digits[i] = d, min
	}

	fmt.Println(digits)
}
