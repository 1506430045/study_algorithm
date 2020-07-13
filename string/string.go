package string

import (
	"strings"
)

func strStr(haystack string, needle string) int {

	return 0
}

func twoSum(numbers []int, target int) []int {
	result := make([]int, 2)

	length := len(numbers)
	if length < 2 {
		return result
	}

	left := 0
	right := length - 1
	sum := 0

	for left < right {
		sum = numbers[left] + numbers[right]
		if sum == target {
			return []int{left + 1, right + 1}
		}
		if sum < target {
			left++
		} else {
			right--
		}
	}
	return result
}

//输入: " 1"
//输出: "1"
func reverseWords(s string) string {
	s = strings.TrimSpace(s) //去掉首位空格
	length := len(s) - 1
	i := length

	end := length
	strs := []string{}
	for i >= 0 {
		if i == 0 {
			strs = append(strs, strings.TrimSpace(s[0:end+1]))
		}
		if s[i] == 32 {
			if s[i-1] == 32 {
				i--
				continue
			}
			strs = append(strs, strings.TrimSpace(s[i+1:end+1]))
			end = i
		}
		i--
	}

	return strings.Join(strs, " ")
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	s := strs[0]
	for i := 1; i < len(strs); i++ {
		s = ss(s, strs[i])
	}
	return s
}

func ss(str1, str2 string) string {
	len1 := len(str1)
	len2 := len(str2)

	i := 0
	var build strings.Builder
	for i < len1 && i < len2 {
		if str1[i] == str2[i] {
			build.WriteByte(str1[i])
		} else {
			break
		}
		i++
	}
	return build.String()
}

func commonChars(A []string) []string {
	length := len(A)
	res := make([]string, 0)
	if length == 0 {
		return res
	}

	m1 := make(map[uint8]int)
	m2 := make(map[uint8]int)
	for i := 0; i < len(A[0]); i++ {
		if _, ok := m1[A[0][i]]; ok {
			m1[A[0][i]]++
		} else {
			m1[A[0][i]] = 1
		}
	}
	for i := 1; i < length; i++ {
		for j := 0; j < len(A[i]); j++ {
			if _, ok := m2[A[i][j]]; ok {
				m2[A[i][j]]++
			} else {
				m2[A[i][j]] = 1
			}
		}
		m1 = compare(m1, m2)
		m2 = make(map[uint8]int, 0)
	}
	for v, times := range m1 {
		for i := 0; i < times; i++ {
			res = append(res, string(v))
		}
	}
	return res
}

func compare(m1, m2 map[uint8]int) map[uint8]int {
	for k, v1 := range m1 {
		if v2, ok := m2[k]; ok {
			if v2 < v1 {
				m1[k] = v2
			}
		} else {
			delete(m1, k)
		}
	}
	return m1
}

func reverseLeftWords(s string, n int) string {
	length := len(s)
	if n >= length {
		return s
	}
	builder := strings.Builder{}
	builder.WriteString(s[n:length])
	builder.WriteString(s[0:n])
	return builder.String()
}
