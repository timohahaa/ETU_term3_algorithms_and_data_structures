package timsort

// decided to make it with a custom func cmp(a, b T) int function instead of constraints.Ordered
// cmp(a, b) = 1 (> 0), a > b
// cmp(a, b) = -1 (< 0), a < b
// cmp(a, b) = 0, a == b
// also made the whole module generic, cause I hate my life, I guess :)

const (
    minimumMinRunSize = 64
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

// func to reverse "runs", that are in descending order
func reverse[T any](arr []T, left, right int) {
    i, j := left, right
    for i <= j {
        arr[i], arr[j] = arr[j], arr[i]
        i++
        j--
    }
}

// no need to use *[]T or return []T, because the function does not make make() calls under the hood
func findRunAndInsertionSort[T any](arr []T, left, right, runSize int, cmp func(a, b T) int) {
    // runs can be only in ascending or STRICTLY descending order
    // this flag is used to validate the run using its first two elements
    descendingFlag := false
}


// no need to use *[]T or return []T, because the function does not make make() calls under the hood
func TimSort[T any](arr []T, cmp func(a, b T) int) {
    // if initial array is to small - just sort it using insertion sort
//    if len(arr) <= minimumMinRunSize {
//        insertionSort[T](arr, 0, len(arr) - 1, cmp)
//        return
//    }
}
