package main

func twoSum(nums []int, target int) []int {
	tmp := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if v, ok := tmp[targe-nums[i]]; ok {
			return []int{v, i}
		}
		tmp[val] = i
	}
	return nil
}
