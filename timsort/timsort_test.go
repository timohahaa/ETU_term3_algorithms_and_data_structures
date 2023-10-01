package timsort

import "testing"

func arraysAreEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestInsertionSort(t *testing.T) {
	arr := []int{4, 3, 2, 10, 12, 1, 5, 6}
	need := []int{1, 2, 3, 4, 5, 6, 10, 12}
	insertionSort[int](arr, 0, len(arr)-1, func(a, b int) int {
		if a > b {
			return 1
		}
		if a < b {
			return -1
		}
		return 0
	})
	if !arraysAreEqual(arr, need) {
		t.Errorf("the array was not sorted right.\n need: %v, got: %v", arr, need)
	}
}
