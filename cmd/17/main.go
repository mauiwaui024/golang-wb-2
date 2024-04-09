// nums := []int{-1, 0, 3, 5, 9, 12}
package main

func main() {
	nums := []int{-1, 0, 3, 5, 9, 12}
	BinarySearch(nums, 0)
}
func BinarySearch(nums []int, target int) int {
	//найти середину и конец

	var start, end int = 0, len(nums) - 1
	return BinarySearchHelper(nums, target, start, end)
}

func BinarySearchHelper(nums []int, target int, start int, end int) int {
	if start > end {
		return -1
	}
	mid := (start + end) / 2
	//2
	if nums[mid] == target {
		return mid
	}

	if nums[mid] > target {

		return BinarySearchHelper(nums, target, start, mid-1)
	}
	if nums[mid] < target {

		return BinarySearchHelper(nums, target, mid+1, end)
	}
	return -1
}
