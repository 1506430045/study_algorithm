package recursive

import "fmt"

func printReverse(s string) {
	helper(s, 0)
}

func helper(s string, i int) {
	if s == "" || i >= len(s) {
		return
	}
	helper(s, i+1)
	fmt.Print(string(s[i]))
}

func reverseString(s []byte) {
	length := len(s)
	for i := 0; i < length/2; i++ {
		s[i], s[length-1-i] = s[length-1-i], s[i]
	}
}
