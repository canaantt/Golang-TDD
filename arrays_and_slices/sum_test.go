package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, expected int) {
		t.Helper()
		if got != expected {
			t.Errorf("got '%d' is not '%d'", got, expected)
		}
	}

	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		expected := 15
		assertCorrectMessage(t, expected, got)
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4}
		got := Sum(numbers)
		expected := 10
		assertCorrectMessage(t, expected, got)
	})

}

func TestSumAll(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, expected []int) {
		t.Helper()
		if !reflect.DeepEqual(got, expected) {
			t.Errorf("got '%v' is not '%v'", got, expected)
		}
	}

	t.Run("Check the subsums of each slice", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 7}, []int{3, 5, 9})
		expected := []int{3, 7, 17}
		assertCorrectMessage(t, got, expected)
	})
}

func TestSumAllTails(t *testing.T) {
	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		expected := []int{2, 9}
		if !reflect.DeepEqual(got, expected) {
			t.Errorf("got %v expected %v", got, expected)
		}
	})
	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		expected := []int{0, 9}
		if !reflect.DeepEqual(got, expected) {
			t.Errorf("got %v expected %v", got, expected)
		}
	})
}
