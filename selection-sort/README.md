## Selection Sort
The selection sort algorithm sorts an array by repeatedly finding the minimum element (considering ascending order) from unsorted part and putting it at the beginning. The algorithm maintains two subarrays in a given array.
- The subarray which is already sorted. 
- Remaining subarray which is unsorted.
 
In every iteration of selection sort, the minimum element (considering ascending order) from the unsorted subarray is picked and moved to the sorted subarray.

**Time Complexity:** O(N<sup>2</sup>)

**Example (Go):**
```
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
```

**Output**:\
``Sorted array: [10 15 25 35 45 55 65]``
