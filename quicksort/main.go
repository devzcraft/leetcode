package main

import "fmt"

func main() {
	arr := []int{9, 4, 6, 7, 1, 2, 9, 0, 3, 6}
	startQuickSort(arr)
	fmt.Println(arr)
}

func startQuickSort(arr []int) {
	quickSort(arr, 0, len(arr)-1)
}

func quickSort(arr []int, low, high int) {
	if low < high {
		p := partition(arr, low, high)
		quickSort(arr, low, p)
		quickSort(arr, p+1, high)
	}
}

func partition(arr []int, low, high int) int {
	pivot := arr[low]
	i := low - 1
	j := high + 1

	for {
		for {
			i++
			if arr[i] >= pivot {
				break
			}
		}

		for {
			j--
			if arr[j] <= pivot {
				break
			}
		}

		if i >= j {
			return j
		}

		arr[i], arr[j] = arr[j], arr[i]
	}
}
