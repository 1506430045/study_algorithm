package sort

func QuickSort(arr []int, left, right int) {
	if left >= right {
		return
	}
	var i, j, base int
	i = left
	j = right
	base = arr[left] //取最左边的数为基准数
	for i < j {
		for arr[j] >= base && i < j {
			j--
		}
		for arr[i] <= base && i < j {
			i++
		}
		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	//基准数归位
	arr[left] = arr[i]
	arr[i] = base
	QuickSort(arr, left, i-1)
	QuickSort(arr, i+1, right)
}
