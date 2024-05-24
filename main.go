package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	data := make([]int, 100)
	for i := 0; i < 100; i++ {
		data[i] = rand.Intn(100)
	}

	ret := SelectionSort(data)
	fmt.Print(ret)
}

func HeapSort(data []int) []int {
	var b []int
	b = append(b, data[0])

	for i := 1; i < len(data); i++ {
		b = UpHeap(b, data[i])
	}

	for i := 0; i < len(data); i++ {
		data[len(data)-1-i] = b[0]
		b = DownHeap(b)
	}

	return data
}

func UpHeap(a []int, n int) []int {
	a = append(a, n)
	child := len(a) - 1
	parent := (child+1)/2 - 1

	for {
		if a[child] > a[parent] {
			a[child], a[parent], child = a[parent], a[child], parent
			parent = (child+1)/2 - 1
			if parent < 0 {
				break
			}
		} else {
			break
		}
	}
	return a
}

func DownHeap(a []int) []int {
	a[0], a[len(a)-1] = a[len(a)-1], a[0]
	a = a[:len(a)-1]

	parent := 0
	var child int

	for {
		child = 2*parent + 1

		// 子どもがいなければ親が葉になるので終了
		if child > len(a)-1 {
			break
		}

		// 親との比較は子どものうち大きい方とのみしたいので前処理
		if child != len(a)-1 {
			if a[child] < a[child+1] {
				child++
			}
		}

		if a[parent] >= a[child] {
			break
		}

		a[parent], a[child] = a[child], a[parent]
		parent = child
	}

	return a
}

func PrintBinaryTree(a []int) {
	fmt.Println("--------Binary tree begin--------")
	depth := int(math.Log2(float64(len(a))) + 1)
	cnt := 1
	for i, v := range a {
		for j := 0; j < int(math.Pow(2, float64(depth-cnt))-1); j++ {
			fmt.Printf(" ")
		}
		fmt.Printf("%d", v)
		for j := 0; j < int(math.Pow(2, float64(depth-cnt))-1); j++ {
			fmt.Printf(" ")
		}
		fmt.Printf(" ")

		if i == int(math.Pow(2, float64(cnt)))-2 {
			fmt.Println("")
			fmt.Println("")
			cnt++
		}
	}

	fmt.Println("\n---------Binary tree end---------")
}

func MergeSort(data []int) []int {
	if len(data) < 2 {
		return data
	}

	if len(data) == 2 {
		if data[0] > data[1] {
			data[0], data[1] = data[1], data[0]
		}
		return data
	}

	half := len(data) / 2

	left := data[:half]
	right := data[half:]

	left = MergeSort(left)
	right = MergeSort(right)
	i, j := 0, 0

	var ret []int
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			ret = append(ret, left[i])
			i++
		} else {
			ret = append(ret, right[j])
			j++
		}
	}

	ret = append(ret, left[i:]...)
	ret = append(ret, right[j:]...)

	return ret
}

func QuickSort(data []int) []int {
	if len(data) < 2 {
		return data
	}

	pivot := data[0]

	var left []int
	var right []int
	for _, v := range data[1:] {
		if pivot > v {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}

	left = QuickSort(left)
	right = QuickSort(right)

	ret := append(left, pivot)
	ret = append(ret, right...)

	return ret
}

func InsertionSort(data []int) []int {
	n := len(data)
	for i := 1; i < n; i++ {
		if data[i] < data[i-1] {
			j := i
			tmp := data[i]
			for 0 < j && tmp < data[j-1] {
				data[j] = data[j-1]
				j--
			}
			data[j] = tmp
		}
	}
	return data
}

func SelectionSort(data []int) []int {
	length := len(data)
	for i := 0; i < length; i++ {
		minIndex, min := i, data[i]
		for j := i + 1; j < length; j++ {
			if min > data[j] {
				minIndex, min = j, data[j]
			}
		}
		data[i], data[minIndex] = data[minIndex], data[i]
	}

	return data
}

func BubbleSort(data []int) []int {
	length := len(data)
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			if data[i] < data[j] {
				data[i], data[j] = data[j], data[i]
			}
		}
	}

	return data
}
