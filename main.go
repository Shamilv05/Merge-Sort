package main

import (
	"./mergesort"
	"log"
	"math/rand"
	"time"
)

func main() {
	numberOfItems := 1000000
	threshold := 2

	items := rand.Perm(numberOfItems)
	start := time.Now()
	arr := mergesort.MergeSort(items, threshold)
	elapsed := time.Since(start)
	log.Printf("Parallel took %s", elapsed)
	log.Printf("Length of array is %v", len(items))

	for i := 1; i < len(arr) - 1; i++ {
		if arr[i] > arr[i + 1] {
			panic("Error in Merge Sort")
		}
	}
}