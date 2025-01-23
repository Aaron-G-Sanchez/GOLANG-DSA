package main

import "fmt"

func twoSum(nums []int, target int) []int {
	if len(nums) == 2 {
		return []int{0, 1}
	}

	// Collect numbers and their index in a map
	numMap := make(map[int]int)
	// value: index

	var result []int
	// loop through nums
	for idx, num := range nums {
		if _, ok := numMap[num]; !ok {
			numMap[num] = idx
		}
	}

	for idx, num := range nums {

		// subtract current iteration from the target and check val is in the map
		diff := target - num
		// if value is in the map
		if val, ok := numMap[diff]; ok && val != idx {
			result = []int{idx, val}
		}
	}
	// return [index of current iteration, map[val]]
	return result
}

func main() {
	// nums := []int{2, 7, 11, 15}
	nums := []int{2, 4, 11, 3}
	// nums := []int{3, 3}
	fmt.Println(twoSum(nums, 6))
}

// N = len(nums) (Length of nums)
// Time = O(n)
// Space = O(n)
