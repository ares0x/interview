package array

import (
	"container/list"
)

// 1. 枚举左边界和右边界，然后长*宽
func maxAreaBasic(height []int) int {
	var max int
	for i := 0; i < len(height)-1; i++ {
		for j := i + 1; j < len(height); j++ {
			area := (j - i) * minFunc(height[i], height[j])
			max = maxFunc(max, area)
		}
	}
	return max
}

func minFuc(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func maxFunc(i, j int) int {
	if i > j {
		return i
	}
	return j
}

// 2.从左右两遍向中间收敛（高度不及边界，面积肯定小）
func maxArea(height []int) int {
	var (
		max int
		i   = 0
		j   = len(height) - 1
	)
	for i < j {
		var minHeight int
		if a[i] < a[j] {
			minHeight = a[i]
			i++
		} else {
			minHeight = a[j]
			j--
		}
		area := (j - i + 1) * minHeight
		max = maxFunc(area, max)
	}
	return max
}
