package timsort

import (
	"math/rand"
	"testing"
)

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
		t.Errorf("the array was not sorted right.\n need: %v\n got: %v", arr, need)
	}
}

func TestBSLowerBound(t *testing.T) {
	cmp := func(a, b int) int {
		if a > b {
			return 1
		}
		if a < b {
			return -1
		}
		return 0
	}
	arr := []int{-7, -4, 3, 4, 9, 9, 12}
	expect := 4
	got := binarySearchLowerBound[int](arr, cmp, 9)
	if got != expect {
		t.Errorf("did not find expected index. Got: %v, expected: %v", got, expect)
	}
}

func TestTimSort(t *testing.T) {
	arr := rand.Perm(1000)
	need := make([]int, 0, 1000)
	for i := range arr {
		need = append(need, i)
	}
	TimSort[int](arr, func(a, b int) int {
		if a < b {
			return -1
		}
		if a > b {
			return 1
		}
		return 0
	})
	if !arraysAreEqual(arr, need) {
		t.Errorf("the array was not sorted right.\n need: %v\n got: %v", arr, need)
	}
}
