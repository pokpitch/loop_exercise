package loop

import "testing"

func TestSumOfFirst(t *testing.T) {
	given := 3
	want := 6

	result := sumOfFirst(given)
	if result != want {
		t.Errorf("sumOfFirst(%d) = %d, want %d", given, result, want)
	}
}
