package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8}
	pos := 3
	nums = append(nums[:pos], nums[pos+1:]...)
	fmt.Println(nums)
}

// func deleteFromSlice() {

// }
