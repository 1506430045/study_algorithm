package heap

import (
	"strings"
)

type element struct {
	Num   int //元素数字
	Times int //数字出现的次数
}

func TopKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if _, ok := m[nums[i]]; ok {
			m[nums[i]]++
		} else {
			m[nums[i]] = 1
		}
	}
	h := make([]element, k) //堆数组
	res := make([]int, 0)   //结果

	for n, v := range m {
		if v <= h[0].Times { //比堆顶还小不处理
			continue
		}
		h[0] = element{
			Num:   n,
			Times: v,
		}
		AdjustHeap(h, 0, k)
	}
	for i := k - 1; i >= 0; i-- {
		res = append(res, h[i].Num)
	}
	return res
}

func AdjustHeap(arr []element, pos int, length int) {
	for {
		//计算孩子位置
		child := pos*2 + 1
		//检查孩子是否越界
		if child >= length {
			break
		}
		//存在右节点且小于左节点
		if child+1 < length && arr[child+1].Times < arr[child].Times {
			child = child + 1
		}

		//检查孩子是否小于于父亲，如果小于则交换
		if arr[pos].Times > arr[child].Times {
			//交换
			arr[pos], arr[child] = arr[child], arr[pos]
			//更新pos，指向子节点
			pos = child
		} else {
			break
		}
	}
}

type element2 struct {
	Key   string
	Times int
}

func topKFrequent(words []string, k int) []string {
	m := make(map[string]int)
	for _, word := range words {
		if _, ok := m[word]; ok {
			m[word]++
		} else {
			m[word] = 1
		}
	}
	h := make([]element2, k) //最小堆初始化
	res := make([]string, 0)
	for key, times := range m {
		if times < h[0].Times {
			continue
		}
		h[0] = element2{
			Key:   key,
			Times: times,
		}
		AdjustHeap2(h, 0, k)
	}
	for i := k - 1; i >= 0; i-- {
		res = append(res, h[i].Key)
	}
	return res
}

func AdjustHeap2(arr []element2, pos int, length int) {
	for {
		//计算孩子位置
		child := pos*2 + 1
		//检查孩子是否越界
		if child >= length {
			break
		}
		//存在右节点且小于左节点
		if child+1 < length && Compare(arr[child], arr[child+1]) == 1 {
			child = child + 1
		}

		//检查孩子是否小于于父亲，如果小于则交换
		if Compare(arr[pos], arr[child]) == 1 {
			//交换
			arr[pos], arr[child] = arr[child], arr[pos]
			//更新pos，指向子节点
			pos = child
		} else {
			break
		}
	}
}

func Compare(e1, e2 element2) int {
	if e1.Times > e2.Times {
		return 1
	} else if e1.Times < e2.Times {
		return -1
	} else {
		return strings.Compare(e1.Key, e2.Key)
	}
}

func smallestK(arr []int, k int) []int {
	BuildHeap(arr[0:k])
	for i := k; i < len(arr); i++ {
		if arr[i] > arr[0] {
			continue
		}
		arr[0] = arr[i]
		AdjustHeap3(arr, 0, k)
	}
	return arr[0:k]
}

func BuildHeap(arr []int) {
	//从底层向上层构建，len(a)/2-1是第一个父节点
	for i := len(arr)/2 - 1; i >= 0; i-- {
		AdjustHeap3(arr, i, len(arr))
	}
}

func AdjustHeap3(arr []int, pos int, length int) {
	for {
		//计算孩子位置
		child := pos*2 + 1
		//检查孩子是否越界
		if child >= length {
			break
		}
		//存在右节点且大于左节点
		if child+1 < length && arr[child+1] > arr[child] {
			child = child + 1
		}

		//检查孩子是否大于父亲，如果大于则交换
		if arr[pos] < arr[child] {
			//交换
			arr[pos], arr[child] = arr[child], arr[pos]
			//更新pos，指向子节点
			pos = child
		} else {
			break
		}
	}
}

func BuildHeap2(arr []int) {
	//从底层向上层构建，len(a)/2-1是第一个父节点
	for i := len(arr)/2 - 1; i >= 0; i-- {
		AdjustHeap4(arr, i, len(arr))
	}
}

func AdjustHeap4(arr []int, pos int, length int) {
	for {
		//计算孩子位置
		child := pos*2 + 1
		//检查孩子是否越界
		if child >= length {
			break
		}
		if child+1 < length && arr[child+1] < arr[child] {
			child = child + 1
		}

		if arr[pos] > arr[child] {
			//交换
			arr[pos], arr[child] = arr[child], arr[pos]
			//更新pos，指向子节点
			pos = child
		} else {
			break
		}
	}
}

func findKthLargest(nums []int, k int) int {
	BuildHeap2(nums[0:k])
	for i := k; i < len(nums); i++ {
		if nums[i] < nums[0] {
			continue
		}
		nums[0] = nums[i]
		AdjustHeap4(nums[0:k], 0, k)
	}
	return nums[0]
}