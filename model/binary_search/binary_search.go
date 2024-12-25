package binary_search

// nums需要从小到大排序，找到第一个不小于target的数字
func LowerBound(nums []int, target int) int {
	l := 0
	r := len(nums) - 1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return nums[l]
}

// nums需要从小到大排序，找到第一个大于target的数字
func UpperBound(nums []int, target int) int {
	l := 0
	r := len(nums) - 1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] <= target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return nums[l]
}
