package other

import (
	"strings"
)

func findRepeatNumber(nums []int) int {
	m := make(map[int]bool, 0)
	for i := 0; i < len(nums); i++ {
		if _, ok := m[nums[i]]; ok {
			return nums[i]
		}
		m[nums[i]] = true
	}
	return -1
}

func findNumberIn2DArray(matrix [][]int, target int) bool {
	m := len(matrix)
	if m == 0 {
		return false
	}
	n := len(matrix[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] > target {
				n = j
			}
			if matrix[i][j] == target {
				return true
			}
		}
	}
	return false
}

func replaceSpace(s string) string {
	build := strings.Builder{}
	for i := 0; i < len(s); i++ {
		if s[i] == 32 {
			build.WriteString("%20")
		} else {
			build.WriteByte(s[i])
		}
	}
	return build.String()
}

func maxProfit(prices []int) int {
	length := len(prices)
	if length < 2 {
		return 0
	}
	minPrice := prices[0]
	max := 0
	for i := 1; i < length; i++ {
		if prices[i]-minPrice > max {
			max = prices[i] - minPrice
		}
		if prices[i] < minPrice {
			minPrice = prices[i]
		}
	}
	return max
}

func singleNumber(nums []int) int {
	for i := 1; i < len(nums); i++ {
		nums[0] ^= nums[i]
	}
	return nums[0]
}

type Stack struct {
	data     []interface{}
	length   int
	capacity int
}

func InitStack() *Stack {
	return &Stack{
		data:     make([]interface{}, 8),
		length:   0,
		capacity: 8,
	}
}

func (s *Stack) Push(val interface{}) {
	if s.length+1 > s.capacity { //需要扩容
		s.capacity = s.capacity * 2
		t := s.data
		s.data = make([]interface{}, s.capacity)
		copy(s.data, t)
	}
	s.data[s.length] = val
	s.length++
}

func (s *Stack) Pop() interface{} {
	if s.Empty() {
		return nil
	}
	data := s.data[s.length-1]
	s.data[s.length-1] = nil
	s.length--
	return data
}

func (s *Stack) Top() interface{} {
	if s.Empty() {
		return nil
	}
	return s.data[s.length-1]
}

func (s *Stack) Empty() bool {
	return s.length == 0
}
