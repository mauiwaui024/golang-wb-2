package main

import "fmt"

func quickSort(arr []int) {
	if len(arr) < 2 {
		return
	}

	pivotIndex := len(arr) / 2
	pivot := arr[pivotIndex]

	left := 0
	right := len(arr) - 1

	arr[pivotIndex], arr[right] = arr[right], arr[pivotIndex]

	for i := range arr {
		if arr[i] < pivot {
			arr[i], arr[left] = arr[left], arr[i]
			left++
		}
	}

	arr[left], arr[right] = arr[right], arr[left]

	quickSort(arr[:left])
	quickSort(arr[left+1:])
}

func main() {
	arr := []int{9, -3, 5, 2, 6, 8, -6, 1, 3}
	fmt.Println("Before sorting:", arr)
	quickSort(arr)
	fmt.Println("After sorting:", arr)
}
