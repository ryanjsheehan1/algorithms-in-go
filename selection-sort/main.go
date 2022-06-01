package main

import "fmt"

func selectionSort(array []int) []int {
	for i := 0; i < len(array)-1; i++ {
		lowestNumberIndex := i
		for j := i + 1; j < len(array); j++ {
			if array[j] < array[lowestNumberIndex] {
				lowestNumberIndex = j
			}
		}
		if lowestNumberIndex != i {
			temp := array[i]
			array[i] = array[lowestNumberIndex]
			array[lowestNumberIndex] = temp
		}
	}
	return array
}

func main() {
	list := []int{65, 55, 45, 35, 25, 15, 10}
	result := selectionSort(list)

	fmt.Printf("Sorted array: %d\n", result)
}
