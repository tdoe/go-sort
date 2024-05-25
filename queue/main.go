package main

import "fmt"

type precess struct {
	name   string
	exceed int
}

type Queue []*precess

func main() {
	var queue Queue
	tmp := Queue{
		&precess{
			name:   "p1",
			exceed: 150,
		},
		&precess{
			name:   "p2",
			exceed: 80,
		},
		&precess{
			name:   "p3",
			exceed: 200,
		},
		&precess{
			name:   "p4",
			exceed: 350,
		},
		&precess{
			name:   "51",
			exceed: 20,
		},
	}
	queue = append(queue, tmp...)

	totalExceed := 0
	for len(queue) > 0 {
		q := queue[0]
		if q.exceed > 100 {
			totalExceed += 100
			q.exceed -= 100
			queue = append(queue[:0], queue[0+1:]...)
			queue = append(queue, q)
		} else {
			totalExceed += q.exceed
			queue = append(queue[:0], queue[0+1:]...)
			fmt.Println(q.name, " ", totalExceed)
		}
	}
}
