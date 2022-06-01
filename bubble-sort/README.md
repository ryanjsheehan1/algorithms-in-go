## Bubble Sort
Bubble Sort is the simplest sorting algorithm that works by repeatedly swapping the adjacent elements if they are in the wrong order. This algorithm is not suitable for large data sets as its average and worst case time complexity is quite high.\
\
**Time Complexity:** O(N<sup>2</sup>)

**Example (Go):**
```
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
```

**Output**:\
``Sorted array: [10 15 25 35 45 55 65]``
