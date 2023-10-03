package timsort

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/timohahaa/ETU_term3_algorithms_and_data_structures/lds"
)

// decided to make it with a custom func cmp(a, b T) int function instead of constraints.Ordered
// cmp(a, b) = 1 (> 0), a > b
// cmp(a, b) = -1 (< 0), a < b
// cmp(a, b) = 0, a == b
// also made the whole module generic, cause I hate my life, I guess :)

// NOTE:
// so basically there are two timsort implementations - the original one - with run finding, stack that stores the runs,
// and merging space overhead
// and the "lazy" one - with fixed run size, no reversing and no stack, it also is O(1) in terms of memory (not counting the stack)
// I made the FULL one, not the lazy one

const (
	minimumMinRunSize = 64
	gallopingParam    = 7
)

func calcMinRun(size int) int {
	temp := 0
	for size >= minimumMinRunSize {
		temp |= (size & 1)
		size >>= 1
	}
	return size + temp
}

// no need to use *[]T or return []T, because the function does not make make() calls under the hood
func insertionSort[T any](arr []T, left, right int, cmp func(a, b T) int) {
	for i := left + 1; i < right+1; i++ {
		valueToInsert := arr[i]
		for j := i - 1; j >= left; j-- {
			if cmp(arr[j], valueToInsert) < 0 {
				// arr[j] < valueToInsert
				break
			}
			// move the element
			arr[j+1] = arr[j]
			arr[j] = valueToInsert
		}
	}
}

func binarySearchLowerBound[T any](arr []T, cmp func(a, b T) int, target T) int {
	low, high := 0, len(arr)
	for low < high {
		mid := (low + high) / 2
		if cmp(arr[mid], target) >= 0 {
			high = mid
		} else {
			low = mid + 1
		}
	}
	if low < len(arr) { // && cmp(arr[low], target) == 0 {
		return low
	}
	return -1
}

// merges two sorted slices into one
// uses galloping for perfomance boost
func merge[T any](left, right []T, cmp func(a, b T) int) []T {
	//fmt.Println()
	//fmt.Println(left, right)
	//fmt.Println()
	merged := make([]T, 0, len(left)+len(right))
	l, r := 0, 0
	countGallop := 0
	for l != len(left) && r != len(right) {
		if countGallop == gallopingParam {
			idx := binarySearchLowerBound[T](left, cmp, right[r])
			if idx == -1 {
				for l != len(left) {
					merged = append(merged, left[l])
					l++
				}
			} else {
				for l != idx {
					merged = append(merged, left[l])
					l++
				}
			}
			countGallop = 0
		} else if countGallop == -gallopingParam {
			idx := binarySearchLowerBound[T](right, cmp, left[l])
			if idx == -1 {
				for r != len(right) {
					merged = append(merged, right[r])
					r++
				}
			} else {
				for r != idx {
					merged = append(merged, right[r])
					r++
				}
			}
			countGallop = 0
		} else if cmp(left[l], right[r]) < 0 {
			merged = append(merged, left[l])
			l++
			if countGallop >= 0 {
				countGallop++
			} else {
				countGallop = 0
			}
		} else {
			merged = append(merged, right[r])
			r++
			if countGallop <= 0 {
				countGallop--
			} else {
				countGallop = 0
			}
		}
	}
	for l != len(left) {
		merged = append(merged, left[l])
		l++
	}
	for r != len(right) {
		merged = append(merged, right[r])
		r++
	}
	return merged
}

// func to reverse "runs", that are in descending order
func reverse[T any](arr []T) {
	i, j := 0, len(arr)-1
	for i <= j {
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
}

// checks if the merge should be permormed and performs it
func mergeIfNeeded[T any](runs *lds.Stack[[]T], cmp func(a, b T) int) {
	if runs.Len() < 2 {
		return
	}
	if runs.Len() == 2 {
		merged := merge[T](runs.Pop(), runs.Pop(), cmp)
		runs.Push(merged)
		return
	}
	// else runs.Len() >= 3
	X := runs.Pop()
	Y := runs.Pop()
	Z := runs.Pop()
	// reorders X, Y and Z in a right order
	reorderXYZ := func() {
		if len(X) > len(Z) {
			X, Z = Z, X
		}
		if len(Y) > len(Z) {
			Y, Z = Z, Y
		}
		if len(Y) < len(X) {
			Y, X = X, Y
		}
	}
	for !runs.Empty() && (len(Z) < len(X)+len(Y) || len(Y) < len(X)) {
		if len(X) < len(Z) {
			Y = merge[T](X, Y, cmp)
			if !runs.Empty() {
				X = runs.Pop()
			} else {
				X = nil
			}
		} else {
			Y = merge[T](Z, Y, cmp)
			if !runs.Empty() {
				Z = runs.Pop()
			} else {
				Z = nil
			}
		}
		reorderXYZ()
	}
	// from biggest to largest
	if Z != nil {
		runs.Push(Z)
	}
	runs.Push(Y)
	if X != nil {
		runs.Push(X)
	}
}

// no need to use *[]T or return []T, because the function does not make make() calls under the hood
func TimSort[T any](arr []T, cmp func(a, b T) int) {
	if len(arr) <= minimumMinRunSize {
		insertionSort[T](arr, 0, len(arr)-1, cmp)
		return
	}

	minRun := calcMinRun(len(arr))
	descendingFlag := false
	// helper function to set descendingFlag
	setFlag := func(i int) {
		if cmp(arr[i], arr[i-1]) < 0 {
			descendingFlag = true
		} else {
			descendingFlag = false
		}
	}
	// find runs, add them to the stack, merge them if neccesary
	runs := lds.NewStack[[]T]()
	run := make([]T, 0, minRun)
	i := 2
	run = append(run, arr[0], arr[1])
	setFlag(1)
	for i != len(arr) {
		if descendingFlag {
			for i != len(arr) && cmp(arr[i], arr[i-1]) < 0 {
				run = append(run, arr[i])
				i++
			}
			// because the order is descending, reverse the run
			reverse[T](run)
			// if there are too little elements in a run, add them until there are minRun elements
			for i != len(arr) && len(run) < minRun {
				run = append(run, arr[i])
				i++
			}
		} else {
			for i != len(arr) && cmp(arr[i], arr[i-1]) >= 0 {
				run = append(run, arr[i])
				i++
			}
			// if there are too little elements in a run, add them until there are minRun elements
			for i != len(arr) && len(run) < minRun {
				run = append(run, arr[i])
				i++
			}
		}
		// sort and push to stack
		insertionSort[T](run, 0, len(run)-1, cmp)
		runs.Push(run)
		// realloc the run slice
		run = make([]T, 0, minRun)
		if i <= len(arr)-2 {
			run = append(run, arr[i])
			i++
			run = append(run, arr[i])
			setFlag(i)
			i++
		} else if i == len(arr)-1 {
			runs.Push([]T{arr[len(arr)-i]})
			i++
		}
		// now merge if needed
		mergeIfNeeded[T](runs, cmp)
	}
	sorted := runs.Pop()
	for i := range sorted {
		arr[i] = sorted[i]
	}
}

// functions test the sorting (show to teacher :))
func equal(a, b []int) bool {
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

func Test(size int) {
	arr := rand.Perm(size)
	need := make([]int, 0, size)
	for i := range arr {
		need = append(need, i)
	}
	fmt.Printf("testing timsort on an array of size %v\n", size)
	//	fmt.Println("before:")
	//	fmt.Printf("arr: %v\nneed: %v\n", arr, need)
	before := time.Now()
	TimSort[int](arr, func(a, b int) int {
		if a < b {
			return -1
		}
		if a > b {
			return 1
		}
		return 0
	})
	after := time.Now()
	if !equal(arr, need) {
		fmt.Println("didnt pass the test")
	} else {
		fmt.Println("test passed")
	}
	//	fmt.Println("after:")
	//	fmt.Printf("arr: %v\nneed: %v\n", arr, need)
	fmt.Printf("time to sort an array of size %v: %v\n", size, time.Duration(after.Sub(before)))
}
