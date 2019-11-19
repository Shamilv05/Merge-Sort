package mergesort

import (
	"sync"
)

func MergeSort(list []int, numberOfWorkers int) []int {

	useThreshold := !(numberOfWorkers <= 1)

	size := len(list)
	middle := size / 2

	if size <= 1 {
		return list
	}

	var left, right []int

	sortInNewRoutine := size > numberOfWorkers && useThreshold

	if !sortInNewRoutine {
		left = MergeSort(list[:middle], numberOfWorkers)
		right = MergeSort(list[middle:], numberOfWorkers)
	} else {
		var wg sync.WaitGroup
		wg.Add(1)

		//worker
		go func() {
			defer wg.Done()
			left = MergeSort(list[:middle], numberOfWorkers - 1)
		}()

		right = MergeSort(list[middle:], numberOfWorkers - 1)

		wg.Wait()
	}

	return merge(left, right)
}

func merge(left, right []int) (result []int) {
	result = make([]int, len(left) + len(right))

	i := 0
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = right[0]
			right = right[1:]
		}
		i++
	}

	for j := 0; j < len(left); j++ {
		result[i] = left[j]
		i++
	}
	for j := 0; j < len(right); j++ {
		result[i] = right[j]
		i++
	}

	return
}