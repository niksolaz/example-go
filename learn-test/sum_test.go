package main

import (
	"testing"
)

func TestArrayOfSum(t *testing.T) {
	numbers := []int8{1, 2, 3, 4, 5}
	result, err := Sum(numbers)
	expected := int8(15)

	if err != nil {
		t.Error("Expected nil but got an error")
	}

	if result != expected {
		t.Errorf("Expected %d but got %d", expected, result)
	}
}

func TestSumInvalidTypesOneNumbers(t *testing.T) {
	numbers := []int8{1}
	_, err := Sum(numbers)
	if err == nil {
		t.Error("Expected an error but got nil. Provide more than 1 number")
	}
}

func TestSumInvalidTypesZeroNumbers(t *testing.T) {
	numbers := []int8{}
	_, err := Sum(numbers)
	if err == nil {
		t.Error("Expected an error but got nil. Provide more than 1 number")
	}
}
