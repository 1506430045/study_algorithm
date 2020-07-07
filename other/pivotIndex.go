package other

import "fmt"

//-1, 1, -2, 2, -4, 2
func pivotIndex(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	var sum1, sum2, i, j int
	i = 1
	j = len(nums) - 2

	sum1 = nums[0]
	sum2 = nums[len(nums)-1]

	for i < j {
		if sum1+nums[i]-sum2 <= sum1-(sum2+nums[j]) {
			sum1 += nums[i]
			i++
		} else {
			sum2 += nums[j]
			j--
		}
	}
	fmt.Println(sum1, sum2)
	if sum1 == sum2 {
		return i
	}
	return -1
}

func searchInsert(nums []int, target int) int {
	length := len(nums)
	if length == 0 || target <= nums[0] {
		return 0
	}

	for i := 1; i < length; i++ {
		if target >= nums[i-1] && target <= nums[i] {
			if i == 0 {
				return 0
			}
			return i
		}
	}
	return length
}
