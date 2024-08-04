package main

import (
	"errors"
	"fmt"
)

func Sum(numbers []int8) (int8, error) {
	if len(numbers) < 2 {
		return 0, errors.New("provide more than 1 number")
	}

	var sumTotal int8
	for _, number := range numbers {
		sumTotal += number
	}

	return sumTotal, nil
}

func main() {
	numbers := []int8{1, 2, 3, 4, 5}
	result, err := Sum(numbers)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Result: ", result)
}
