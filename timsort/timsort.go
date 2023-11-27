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
func insertionSort[T any](arr *lds.DynamicArray[T], left, right int, cmp func(a, b T) int) {
	for i := left + 1; i < right+1; i++ {
		valueToInsert := arr.GetNoError(i)
		for j := i - 1; j >= left; j-- {
			if cmp(arr.GetNoError(j), valueToInsert) < 0 {
				// arr[j] < valueToInsert
				break
			}
			// move the element
			arr.Set(j+1, arr.GetNoError(j))
			arr.Set(j, valueToInsert)
		}
	}
}

func binarySearchLowerBound[T any](arr *lds.DynamicArray[T], cmp func(a, b T) int, target T) int {
	low, high := 0, arr.Len()
	for low < high {
		mid := (low + high) / 2
		if cmp(arr.GetNoError(mid), target) >= 0 {
			high = mid
		} else {
			low = mid + 1
		}
	}
	if low < arr.Len() { // && cmp(arr[low], target) == 0 {
		return low
	}
	return -1
}

// merges two sorted slices into one
// uses galloping for perfomance boost
func merge[T any](left, right *lds.DynamicArray[T], cmp func(a, b T) int) *lds.DynamicArray[T] {
	merged := lds.NewDArray[T]()
	//merged.EnsureCapacity(left.Len() + right.Len())
	l, r := 0, 0
	countGallop := 0
	for l != left.Len() && r != right.Len() {
		if countGallop == gallopingParam {
			idx := binarySearchLowerBound[T](left, cmp, right.GetNoError(r))
			if idx == -1 {
				for l != left.Len() {
					merged.PushBack(left.GetNoError(l))
					l++
				}
			} else {
				for l != idx {
					merged.PushBack(left.GetNoError(l))
					l++
				}
			}
			countGallop = 0
		} else if countGallop == -gallopingParam {
			idx := binarySearchLowerBound[T](right, cmp, left.GetNoError(l))
			if idx == -1 {
				for r != right.Len() {
					merged.PushBack(right.GetNoError(r))
					r++
				}
			} else {
				for r != idx {
					merged.PushBack(right.GetNoError(r))
					r++
				}
			}
			countGallop = 0
		} else if cmp(left.GetNoError(l), right.GetNoError(r)) < 0 {
			merged.PushBack(left.GetNoError(l))
			l++
			if countGallop >= 0 {
				countGallop++
			} else {
				countGallop = 0
			}
		} else {
			merged.PushBack(right.GetNoError(r))
			r++
			if countGallop <= 0 {
				countGallop--
			} else {
				countGallop = 0
			}
		}
	}
	for l != left.Len() {
		merged.PushBack(left.GetNoError(l))
		l++
	}
	for r != right.Len() {
		merged.PushBack(right.GetNoError(r))
		r++
	}
	return merged
}

// func to reverse "runs", that are in descending order
func reverse[T any](arr *lds.DynamicArray[T]) {
	i, j := 0, arr.Len()-1
	for i <= j {
		tI := arr.GetNoError(i)
		tJ := arr.GetNoError(j)
		arr.Set(i, tJ)
		arr.Set(j, tI)
		i++
		j--
	}
}

// checks if the merge should be permormed and performs it
func mergeIfNeeded[T any](runs *lds.Stack[*lds.DynamicArray[T]], cmp func(a, b T) int) {
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
		if X.Len() > Z.Len() {
			X, Z = Z, X
		}
		if Y.Len() > Z.Len() {
			Y, Z = Z, Y
		}
		if Y.Len() < X.Len() {
			Y, X = X, Y
		}
	}
	var XisNil, ZisNil = false, false
	for !runs.Empty() && (Z.Len() < X.Len()+Y.Len() || Y.Len() < X.Len()) {
		if X.Len() < Z.Len() {
			Y = merge[T](X, Y, cmp)
			if !runs.Empty() {
				X = runs.Pop()
			} else {
				XisNil = true
			}
		} else {
			Y = merge[T](Z, Y, cmp)
			if !runs.Empty() {
				Z = runs.Pop()
			} else {
				ZisNil = true
			}
		}
		reorderXYZ()
	}
	// from biggest to largest
	if !ZisNil {
		runs.Push(Z)
	}
	runs.Push(Y)
	if !XisNil {
		runs.Push(X)
	}
}

// no need to use *[]T or return []T, because the function does not make make() calls under the hood
func TimSort[T any](arr *lds.DynamicArray[T], cmp func(a, b T) int) {
	if arr.Len() <= minimumMinRunSize {
		insertionSort[T](arr, 0, arr.Len()-1, cmp)
		return
	}

	minRun := calcMinRun(arr.Len())
	descendingFlag := false
	// helper function to set descendingFlag
	setFlag := func(i int) {
		if cmp(arr.GetNoError(i), arr.GetNoError(i-1)) < 0 {
			descendingFlag = true
		} else {
			descendingFlag = false
		}
	}
	// find runs, add them to the stack, merge them if neccesary
	runs := lds.NewStack[*lds.DynamicArray[T]]()
	run := lds.NewDArray[T]()
	//run.EnsureCapacity(minRun)
	i := 2
	run.PushBack(arr.GetNoError(0))
	run.PushBack(arr.GetNoError(1))
	setFlag(1)
	for i != arr.Len() {
		if descendingFlag {
			for i != arr.Len() && cmp(arr.GetNoError(i), arr.GetNoError(i-1)) < 0 {
				run.PushBack(arr.GetNoError(i))
				i++
			}
			// because the order is descending, reverse the run
			reverse[T](run)
			// if there are too little elements in a run, add them until there are minRun elements
			for i != arr.Len() && run.Len() < minRun {
				run.PushBack(arr.GetNoError(i))
				i++
			}
		} else {
			for i != arr.Len() && cmp(arr.GetNoError(i), arr.GetNoError(i-1)) >= 0 {
				run.PushBack(arr.GetNoError(i))
				i++
			}
			// if there are too little elements in a run, add them until there are minRun elements
			for i != arr.Len() && run.Len() < minRun {
				run.PushBack(arr.GetNoError(i))
				i++
			}
		}
		// sort and push to stack
		insertionSort[T](run, 0, run.Len()-1, cmp)
		runs.Push(run)
		// realloc the run slice
		run = lds.NewDArray[T]()
		//		run.EnsureCapacity(minRun)
		if i <= arr.Len()-2 {
			run.PushBack(arr.GetNoError(i))
			i++
			run.PushBack(arr.GetNoError(i))
			setFlag(i)
			i++
		} else if i == arr.Len()-1 {
			lastRun := lds.NewDArray[T]()
			lastRun.PushBack(arr.GetNoError(arr.Len() - 1))
			runs.Push(lastRun)
			i++
		}
		// now merge if needed
		mergeIfNeeded[T](runs, cmp)
	}
	sorted := runs.Pop()
	for i := 0; i < sorted.Len(); i++ {
		arr.Set(i, sorted.GetNoError(i))
	}
}

// functions test the sorting (show to teacher :))
func Equal(a, b *lds.DynamicArray[int]) bool {
	if a.Len() != b.Len() {
		return false
	}
	for i := 0; i < a.Len(); i++ {
		if a.GetNoError(i) != b.GetNoError(i) {
			return false
		}
	}
	return true
}

func Test(size int) {
	perm := rand.Perm(size)
	arr := lds.NewDArray[int]()
	need := lds.NewDArray[int]()

	for i := range perm {
		need.PushBack(i)
		arr.PushBack(perm[i])
	}
	fmt.Printf("testing timsort on an array of size %v\n", size)
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
	if !Equal(arr, need) {
		fmt.Println("didnt pass the test")
	} else {
		fmt.Println("test passed")
	}
	fmt.Printf("time to sort an array of size %v: %v\n", size, time.Duration(after.Sub(before)))
}
