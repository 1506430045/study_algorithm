package dynamic_program

const (
	IntSize  = 32 << (^uint(0) >> 63)
	UintSize = 32 << (^uint(0) >> 63)
)

const (
	MaxInt  = 1<<(IntSize-1) - 1
	MinInt  = -1 << (IntSize - 1)
	MaxUint = 1<<UintSize - 1
)

func CoinChange(coins []int, amount int) int {
	if amount < 1 {
		return 0
	}
	count := []int{}
	return doCoinChange(coins, amount, count)
}

func doCoinChange(coins []int, rem int, count []int) int {
	if rem < 0 {
		return -1
	}
	if rem == 0 {
		return 0
	}
	if count[rem-1] != 0 {
		return count[rem-1]
	}
	min := MinInt
	for _, coin := range coins {
		res := doCoinChange(coins, rem-coin, count)
		if res >= 0 && res < min {
			min = res + 1
		}
	}
	if min == MinInt {
		count[rem-1] = -1
	} else {
		count[rem-1] = min
	}
	return count[rem-1]
}