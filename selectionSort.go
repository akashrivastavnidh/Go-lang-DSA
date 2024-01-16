package main

import "fmt"

// SelectionSort performs selection sort on a slice of integers
func SelectionSort(arr []int) {
	length := len(arr)

	// Traverse through all array elements
	for i := 0; i < length-1; i++ {
		// Find the minimum element in the unsorted region
		minIndex := i
		for j := i + 1; j < length; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}

		// Swap the found minimum element with the first element
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}

func main() {
	// Example usage of SelectionSort
	arr := []int{64, 25, 12, 22, 11}
	fmt.Println("Unsorted array:", arr)

	SelectionSort(arr)

	fmt.Println("Sorted array:", arr)
}
