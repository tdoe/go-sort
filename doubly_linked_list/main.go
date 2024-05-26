package main

import (
	"fmt"
	"strconv"
	"strings"
)

type linkedList []int

func (l *linkedList) insert(val int) {
	tmp := []int{val}
	*l = append(tmp, *l...)
}

func (l *linkedList) delete(val int) {
	for i, v := range *l {
		if v == val {
			*l = append((*l)[:i], (*l)[i+1:]...)
			break
		}
	}
}

func (l *linkedList) deleteFirst() {
	*l = (*l)[1:]
}

func (l *linkedList) deleteLast() {
	*l = (*l)[:len(*l)-1]
}

func main() {
	args := []string{"insert 5", "insert 2", "insert 3", "insert 1", "delete 3", "insert 6", "delete 5"}

	linkedList := make(linkedList, 0, 7)
	for _, arg := range args {
		commands := strings.Split(arg, " ")
		command := commands[0]
		val, _ := strconv.Atoi(commands[1])

		switch command {
		case "insert":
			linkedList.insert(val)
		case "delete":
			linkedList.delete(val)
		case "deleteFirst":
			linkedList.deleteFirst()
		case "deleteLast":
			linkedList.deleteLast()
		}
	}

	fmt.Println(linkedList)
}
