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
