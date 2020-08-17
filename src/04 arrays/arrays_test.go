package array

import "testing"

// Version 1
func TestSum(t *testing.T) {

	numbers := [5]int{1, 2, 3, 4, 5}

	got := Sum(numbers)
	want := 15

	if want != got {
		t.Errorf("got %d want %d given, %v", got, want, numbers)
	}
}

// Version 2:
func TestSum2(t *testing.T) {

	numbers := [5]int{1, 2, 3, 4, 5}

	got := Sum2(numbers)
	want := 15

	if got != want {
		t.Errorf("got %d want %d given, %v", got, want, numbers)
	}
}

// Version 3:
func TestSum3(t *testing.T) {

	t.Run("collections of any size", func(t *testing.T) {

		numbers := []int{1, 2, 3}

		got := Sum3(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

}
