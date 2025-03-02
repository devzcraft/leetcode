package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(binarysearch(arr, 9))

}

func binarysearch(arr []int, target int) bool {
	low := 0
	high := len(arr) - 1
	for low <= high {
		med := (low + high) / 2
		if arr[med] == target {
			return true
		}

		if arr[med] < target {
			low = med + 1
		}

		if arr[med] > target {
			high = med - 1
		}
	}

	return false
}
