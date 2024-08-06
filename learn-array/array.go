package main

import (
	"fmt"
)

type Arr struct {
	i int
	b bool
}

func rangeExample() {
	var a = [5]int{1, 2, 3, 4, 5}
	for i, v := range a {
		fmt.Println("index:", i, "value:", v)
	}
}

func sliceMake() {
	s := make([]int, 2, 5)
	fmt.Println("s:", s)
	fmt.Println("len(s):", len(s))
	fmt.Println("cap(s):", cap(s))
	s[1] = 99
	s = append(s, 3)
	fmt.Println("new s:", s)
	fmt.Println("len(s):", len(s))
	fmt.Println("cap(s):", cap(s))
}

func sliceNil() {
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
}

func sliceLenCap() {
	s := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("s:", s)
	s = s[3:]
	fmt.Println("slice s[3:]:", s)
	fmt.Println("len(s):", len(s))
	fmt.Println("cap(s):", cap(s))
}

func readStr(idx int, arr []Arr) (int, bool) {
	fmt.Println("i:", arr[idx].i, "b:", arr[idx].b)
	return arr[idx].i, arr[idx].b
}

func sliceLiteral() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []Arr{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
	readStr(3, s)
}

func sliceExample() {
	var s = [5]bool{true, true, false, false, true}
	c := s[0:2]    // slice of s
	fmt.Println(c) // [true true]
}

func arrTwo() {
	var b = [3]string{"You", "are", "Welcome"}
	fmt.Println(b[0], b[1], b[2]) // You are Welcome
}

func arrOne() {
	var a [5]int
	a[0] = 1
	a[1] = 123
	a[2] = 1234
	a[3] = 12345
	a[4] = 123456
	fmt.Println(a)    // [1 123 1234 12345 123456]
	fmt.Println(a[3]) // 12345
}

func main() {
	arrOne()
	arrTwo()

	sliceExample()
	sliceLiteral()

	sliceLenCap()

	sliceNil()

	sliceMake()

	rangeExample()
}
