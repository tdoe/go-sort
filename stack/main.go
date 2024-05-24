package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	args := "1 2 + 3 4 - *"

	stack := strings.Split(args, " ")

	cursol := 0
	for len(stack) > 1 {
		if stack[cursol] == "+" {
			a, _ := strconv.Atoi(stack[cursol-2])
			b, _ := strconv.Atoi(stack[cursol-1])
			tmp := append(stack[:cursol-2], strconv.Itoa(a+b))
			stack = append(tmp, stack[cursol+1:]...)
			cursol -= 1
		} else if stack[cursol] == "-" {
			a, _ := strconv.Atoi(stack[cursol-2])
			b, _ := strconv.Atoi(stack[cursol-1])
			tmp := append(stack[:cursol-2], strconv.Itoa(a-b))
			stack = append(tmp, stack[cursol+1:]...)
			cursol -= 1
		} else if stack[cursol] == "*" {
			a, _ := strconv.Atoi(stack[cursol-2])
			b, _ := strconv.Atoi(stack[cursol-1])
			tmp := append(stack[:cursol-2], strconv.Itoa(a*b))
			stack = append(tmp, stack[cursol+1:]...)
			cursol -= 1
		} else {
			cursol++
		}
	}
	fmt.Println(stack[0])
}
