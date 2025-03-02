package main

import (
	"fmt"
	"sort"
)

func main() {
	input := []int{-1, 0, 1, 2, -1, -4}
	result := threeSum(input)
	fmt.Println(result)
}

func threeSum(nums []int) [][]int {
	result := [][]int{}
	sum_map := make(map[[3]int]struct{})

	length := len(nums)
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		low := i + 1
		high := length - 1
		sum := 0 - nums[i]
		for low < high {
			switch {
			case nums[high]+nums[low] > sum:
				high--
			case nums[low]+nums[high] < sum:
				low++
			case nums[high]+nums[low] == sum:
				temp := [3]int{nums[i], nums[low], nums[high]}
				if _, ok := sum_map[temp]; !ok {
					result = append(result, temp[:])
					sum_map[temp] = struct{}{}
				}
				high--
				low++
			}
		}
	}

	return result
}
