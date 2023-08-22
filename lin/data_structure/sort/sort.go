package sort

// 冒泡排序
// 第一个元素与所有元素比较，最后最大的冒泡上最上面
func bubbleSort(list []int) {
	for i := 0; i < len(list); i++ {
		maxIndex := 0
		for j := 1; j < len(list); j++ {
			if list[j] > list[maxIndex] {
				maxIndex = j
			} else {
				swap(list, j, maxIndex)
			}
		}
	}
}

func swap(list []int, i, j int) {
	tmp := list[i]
	list[i] = list[j]
	list[j] = tmp
}

// 选择排序
// 一轮遍历后选择最大的放到最后
func selectionSort(list []int) {
	for i := 0; i < len(list)-1; i++ {
		maxIndex := 0
		for j := 1; j < len(list)-i; j++ {
			if list[j] > list[maxIndex] {
				maxIndex = j
			}
		}
		swap(list, maxIndex, len(list)-i-1)
	}
}

// 插入排序
// 分前后两个数组，前一个是排好序的，后一个是待排序的
func insertionSort(list []int) {
	start := list[0]
	end := list[0]
	for i := 1; i < len(list); i++ {
		elem := list[i]
		if elem <= start {
			for j := i - 1; j >= 0; j-- {
				if list[j] > elem {
					list[j+1] = list[j]
				}
			}
			list[0] = elem
		} else if elem >= end {
			continue
		} else {
			// 插入到中间
			for j := i - 1; j >= 0; j-- {
				if list[j] > elem {
					list[j+1] = list[j]
				} else {
					list[j+1] = elem
					break
				}
			}
		}
	}
}

// 归并排序
// 每一次递归分成两半，直到每一半长度都小于等于1，再合并
func mergeSort(list []int) {
	var merge func(start, mid, end int)
	merge = func(start, mid, end int) {
		list1 := make([]int, mid-start+1)
		list2 := make([]int, end-mid)
		copy(list1, list[start:mid+1])
		copy(list2, list[mid+1:end+1])
		index := start
		for len(list1) > 0 || len(list2) > 0 {
			if len(list1) == 0 {
				list[index] = list2[0]
				list2 = list2[1:]
				index++
				continue
			}
			if len(list2) == 0 {
				list[index] = list1[0]
				list1 = list1[1:]
				index++
				continue
			}
			if list1[0] <= list2[0] {
				list[index] = list1[0]
				list1 = list1[1:]
				index++
			} else {
				list[index] = list2[0]
				list2 = list2[1:]
				index++
			}
		}
	}

	var split func(start, end int)
	split = func(start, end int) {
		if start >= end {
			return
		}
		mid := (start + end) >> 1
		split(start, mid)
		split(mid+1, end)
		merge(start, mid, end)
	}
	split(0, len(list)-1)
}

// 快速排序
// 每一次都找到中间的元素，然后左右两边再找
// 第一个元素为哨兵，左右两个指针，比哨兵小的放左边，比哨兵大的放右边，右指针先移动
func quickSort(list []int) {
	var tracking func(start, end int)
	tracking = func(start, end int) {
		t1, t2 := start, end
		if start >= end {
			return
		}
		sentinel := start
		start++
		for start < end {
			for start < end && list[end] >= list[sentinel] {
				end--
			}
			for start < end && list[start] <= list[sentinel] {
				start++
			}
			if start != end {
				swap(list, start, end)
			}
		}
		if list[start] < list[sentinel] {
			swap(list, start, sentinel)
		}
		tracking(t1, start-1)
		tracking(start, t2)
	}
	tracking(0, len(list)-1)
}
