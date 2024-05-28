package main

import "fmt"

func solve(n, i, m int, A []int) bool {
	if m == 0 {
		return true
	}

	if i >= n {
		return false
	}

	return solve(n, i+1, m, A) || solve(n, i+1, m-A[i], A)
}

func main() {

	A := []int{1, 5, 7, 10, 21}
	M := []int{2, 4, 17, 8}

	for _, m := range M {
		if solve(len(A), 0, m, A) {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}
