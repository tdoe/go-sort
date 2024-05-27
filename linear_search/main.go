package main

import "fmt"

func linearSearch(digits []int, target int) bool {
	i := 0
	digits = append(digits, target)
	for digits[i] != target {
		i++
	}

	return i != len(digits)
}

func main() {
	digits := []int{1, 2, 3, 4, 5}
	target := []int{3, 4, 1}

	ans := 0
	for _, t := range target {
		if linearSearch(digits, t) {
			ans++
		}
	}

	fmt.Println(ans)
}
