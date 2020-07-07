package other

func generate(numRows int) [][]int {
	var i, j int
	var res = [][]int{}

	for i = 0; i < numRows; i++ {
		res = append(res, []int{})
		for j = 0; j < numRows; j++ {
			if j > i {
				continue
			}
			if j == 0 || i == j {
				res[i] = append(res[i], 1)
			} else {
				res[i] = append(res[i], res[i-1][j-1]+res[i-1][j])
			}
		}
	}
	return res
}

func getRow(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	}
	if rowIndex == 1 {
		return []int{1, 1}
	}
	var i, j int
	var temp []int
	var res = []int{1, 1}

	for i = 2; i <= rowIndex; i++ {
		temp = res
		res = make([]int, i+1)
		for j = 0; j < i+1; j++ {
			if j == 0 || j == i {
				res[j] = 1
			} else {
				res[j] = temp[j-1] + temp[j]
			}
		}
	}
	return res
}

var m = make(map[int]int)

func fib(N int) int {
	if v, ok := m[N]; ok {
		return v
	}
	if N < 2 {
		return N
	} else {
		fn := fib(N-1) + fib(N-2)
		m[N] = fn
		return fn
	}
}

func climbStairs(n int) int {
	m := map[int]int{
		1: 1,
		2: 2,
	}
	if n <= 2 {
		return m[n]
	} else {
		for i := 3; i <= n; i++ {
			m[i] = m[i-1] + m[i-2]
		}
	}
	return m[n]
}
