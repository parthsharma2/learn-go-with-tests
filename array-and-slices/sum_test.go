package main

import "testing"

func TestSum(t *testing.T) {
	assertEqual := func(t *testing.T, got, want int, numbers []int) {
		t.Helper()

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	}

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 5}

		got := Sum(numbers)
		want := 20

		assertEqual(t, got, want, numbers)
	})
}