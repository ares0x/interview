package array

//1.  遍历，统计0的个数，遍历追加0 O(n^2)

//2. 引入新数组

//3. 双指针/快慢指针 调换位置
func MoveZeros(nums []int) {
	l, r, n := 0, 0, len(nums)
	for r < n {
		if nums[r] != 0 {
			nums[l], nums[r] = nums[r], nums[l]
			l++
		}
		r++
	}
}
