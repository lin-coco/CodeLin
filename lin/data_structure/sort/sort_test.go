package sort

import (
	"testing"
)

func TestBubbleSort(t *testing.T) {
	list := []int{9, 8, 7, 6, 5, 4, 3, 2}
	bubbleSort(list)
	for i := 1; i < len(list); i++ {
		if list[i-1] > list[i] {
			t.Errorf("排序失败: %v", list)
			t.Fail()
		}
	}
	t.Log(list)
}

func TestSelectionSort(t *testing.T) {
	list := []int{9, 8, 7, 6, 5, 4, 3, 2}
	selectionSort(list)
	for i := 1; i < len(list); i++ {
		if list[i-1] > list[i] {
			t.Errorf("排序失败: %v", list)
			t.Fail()
			return
		}
	}
	t.Log(list)
}

func TestInsertionSort(t *testing.T) {
	list := []int{9, 8, 7, 6, 5, 4, 3, 2}
	insertionSort(list)
	for i := 1; i < len(list); i++ {
		if list[i-1] > list[i] {
			t.Errorf("排序失败: %v", list)
			t.Fail()
			return
		}
	}
	t.Log(list)
}

func TestMergeSort(t *testing.T) {
	list := []int{9, 8, 7, 6, 5, 4, 3, 2}
	mergeSort(list)
	for i := 1; i < len(list); i++ {
		if list[i-1] > list[i] {
			t.Errorf("排序失败: %v", list)
			t.Fail()
			return
		}
	}
	t.Log(list)
}

func TestQuickSort(t *testing.T) {
	list := []int{9, 8, 7, 6, 5, 4, 3, 2}
	quickSort(list)
	for i := 1; i < len(list); i++ {
		if list[i-1] > list[i] {
			t.Errorf("排序失败: %v", list)
			t.Fail()
			return
		}
	}
	t.Log(list)
}
