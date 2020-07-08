package sort

func AdjustHeap(arr []int, pos int, length int) {
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

func BuildHeap(arr []int) {
	//从底层向上层构建，len(a)/2-1是第一个父节点
	for i := len(arr)/2 - 1; i >= 0; i-- {
		AdjustHeap(arr, i, len(arr))
	}
}

func HeapSort(arr []int) {
	//构建大顶堆
	BuildHeap(arr)

	//首尾交换 重新构建打顶堆
	for i := len(arr) - 1; i >= 0; i-- {
		arr[i], arr[0] = arr[0], arr[i]
		AdjustHeap(arr, 0, i)
	}
}

func getLeastNumbers(arr []int, k int) []int {
	BuildHeap(arr[0:k])
	for i := k; i < len(arr); i++ {
		if arr[i] > arr[0] { //大于堆顶元素 不需要调整
			continue
		}
		arr[0], arr[i] = arr[i], arr[0]
		AdjustHeap(arr[0:k], 0, k)
	}
	return arr[0:k]
}
