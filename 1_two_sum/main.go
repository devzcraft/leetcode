package main

import "fmt"

func main() {
	nums := []int{5, 3, 8, 4, 2, 7, 1, 10}
	target := 9
	result := twoSum(nums, target)

	fmt.Println(result)
}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))

	for i, n := range nums {
		if j, ok := m[target-n]; ok {
			return []int{i, j}
		}

		m[n] = i
	}

	return []int{}
}
