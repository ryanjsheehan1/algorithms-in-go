package main

import "fmt"

func binarySearch(slice []int, searchValue int) int {
	lowerBound := 0
	upperBound := len(slice) - 1

	for lowerBound <= upperBound {

		midpoint := (upperBound + lowerBound) / 2

		midpointValue := slice[midpoint]

		if searchValue == midpointValue {
			return midpoint
		} else if searchValue < midpointValue {
			upperBound = midpoint - 1
		} else if searchValue > midpointValue {
			lowerBound = midpoint + 1
		}
	}
	return -1
}

func main() {
	s := []int{2, 3, 4, 10, 40}
	sv := 10

	result := binarySearch(s, sv)

	if result == -1 {
		fmt.Printf("Element is not present present\n")
	} else {
		fmt.Printf("Element is present at index %d\n", result)
	}
}
