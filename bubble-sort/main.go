package main

import "fmt"

func bubbleSort(list []int) []int {
	n := len(list) - 1
	sorted := false

	for !sorted {
		sorted = true
		for i := range list[:n] {
			if list[i] > list[i+1] {
				list[i], list[i+1] = list[i+1], list[i]
				sorted = false
			}

		}
		n -= 1
	}
	return list
}

func main() {
	list := []int{65, 55, 45, 35, 25, 15, 10}

	result := bubbleSort(list)

	fmt.Printf("Sorted array: %d\n", result)

}
