package array

func Intersect(nums1 []int, nums2 []int) []int {
	res := make([]int, 0)
	m := make(map[int]int)
	for _, v := range nums1 {
		if _, ok := m[v]; ok {
			m[v]++
		} else {
			m[v] = 1
		}
	}
	for _, v := range nums2 {
		if times, ok := m[v]; ok && times > 0 {
			res = append(res, v)
			m[v]--
		}
	}
	return res
}

func minArray(numbers []int) int {
	var left, right, mid int

	left = 0
	right = len(numbers) - 1
	for left < right {
		mid = left + (right-left)/2
		if numbers[mid] == numbers[right] { //等于右边升序序列的最大值 mid可能在zuo边也可能在右边升序序列 3333333451233、33123333333
			right = right - 1
			continue
		}
		if numbers[mid] > numbers[right] { //大于右边升序序列的最大值 所以mid在左边升序序列 最小值在mid右边  34512
			left = mid + 1
		} else { //小于右边升序序列最大值 所以mid在右边升序序列 最小值在mid的左边或者他自己 56712345
			right = mid
		}
	}
	return numbers[left]
}

func majorityElement(nums []int) int {
	m := make(map[int]int)
	var times int
	var ok bool

	length := len(nums)
	for i := 0; i < length; i++ {
		if times, ok = m[nums[i]]; ok {
			m[nums[i]] ++
		} else {
			m[nums[i]] = 1
		}
		if times+1 >= length/2 {
			return nums[i]
		}
	}
	return 0
}

func exchange(nums []int) []int {
	length := len(nums)
	if length <= 1 {
		return nums
	}
	i := 0
	j := length - 1
	for i < j {
		if nums[i]%2 == 1 { //奇数不处理
			i++
			continue
		}
		if nums[j]%2 == 0 {
			j--
			continue
		}
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
	return nums
}

func myPow(x float64, n int) float64 {
	var res float64 = 1
	var flag = false
	if n == 0 {
		return res
	}
	if n < 0 {
		n = -n
		flag = true
	}
	for i := 1; i <= n; i++ {
		res *= x
	}
	if flag {
		return 1 / res
	}
	return res
}

func isStraight(nums []int) bool {
	m := make(map[int]bool)
	min := nums[0]
	max := nums[0]
	for i := 0; i < 5; i++ {
		if nums[i] == 0 {
			continue
		}
		if min == 0 {
			min = nums[i]
		}
		if _, ok := m[nums[i]]; ok {
			return false
		}
		m[nums[i]] = true
		if nums[i] < min {
			min = nums[i]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max-min <= 4
}

func twoSum(nums []int, target int) []int {
	res := make([]int, 0)
	length := len(nums)
	if length < 2 {
		return res
	}
	i := 0
	j := length - 1
	sum := 0
	for i < j {
		sum = nums[i] + nums[j]
		if sum == target {
			res = append(res, nums[i], nums[j])
			return res
		}
		if sum < target {
			i++
		} else {
			j--
		}
	}
	return res
}
