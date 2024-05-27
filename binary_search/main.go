package main

import "fmt"

func binarySearch(digits []int, key int) int {
	left, right := 0, len(digits)-1

	for left < right {
		mid := (left + right) / 2
		if digits[mid] == key {
			return mid
		}
		if digits[mid] > key {
			right = mid
		} else {
			left = mid
		}
	}

	return -1
}

func main() {
	digits := []int{1, 2, 3, 4, 5}
	targets := []int{3, 4, 1}

	ans := 0
	for _, key := range targets {
		index := binarySearch(digits, key)
		if index != -1 {
			ans++
		}
	}
	fmt.Println(ans)
}
