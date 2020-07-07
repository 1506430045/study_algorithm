package search

//给定一个 n 个元素有序的（升序）整型数组 nums 和一个目标值 target  ，写一个函数搜索 nums 中的 target，如果目标值存在返回下标，否则返回 -1。
func search(nums []int, target int) int {
	var length, left, right, mid int

	length = len(nums)

	left = 0
	right = length - 1

	for left <= right {
		mid = left + (right-left)/2
		if target > nums[mid] {
			left = mid + 1
		} else if target < nums[mid] {
			right = mid - 1
		} else {
			return mid
		}
	}
	return -1
}

func mySqrt(x int) int {
	if x < 0 {
		return -1
	}
	var left, right, mid int

	left = 0
	right = x

	for left <= right {
		mid = left + (right-left)/2
		if mid*mid <= x && (mid+1)*(mid+1) > x {
			return mid
		}
		if mid*mid < x {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func search2(nums []int, target int) int {

	return 0
}

func firstBadVersion(n int) int {
	if n <= 0 {
		return 0
	}
	var mid, left, right int
	var cur bool

	left = 0
	right = n

	for left < right {
		mid = left + (right-left)/2

		cur = isBadVersion(mid)
		if cur {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	if mid+1 > n {
		return mid
	}
	return mid + 1
}

func isBadVersion(version int) bool {
	return false
}
