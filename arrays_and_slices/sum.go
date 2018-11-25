package main

import (
	"fmt"
)

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum = sum + number
	}
	return sum
}

func SumAll(slices ...[]int) []int {
	// sumSlice := make([]int, len(slices))
	// for i, slice := range slices {
	// sumSlice[i] = Sum(slice)
	// }
	var sumSlice []int
	for _, slice := range slices {
		sumSlice = append(sumSlice, Sum(slice))
	}
	return sumSlice
}

func SumAllTails(slices ...[]int) []int {
	var sumSlice []int
	for _, slice := range slices {
		if len(slice) != 0 {
			tail := Sum(slice[1:])
			sumSlice = append(sumSlice, tail)
		} else {
			sumSlice = append(sumSlice, 0)
		}
	}
	return sumSlice
}

func main() {
	fmt.Println("There is a sum function in this package.")
}
