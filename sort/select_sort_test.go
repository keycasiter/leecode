package sort

import (
	"fmt"
	"testing"
)

func SelectSort(arr []int) []int {
	cnt := 0
	for i := 0; i < len(arr); i++ {
		minIndex := i
		for j := i; j < len(arr); j++ {
			if arr[minIndex] > arr[j] {
				minIndex = j
			}
			cnt++
		}
		temp := arr[minIndex]
		arr[minIndex] = arr[i]
		arr[i] = temp
	}
	fmt.Println("cnt -> ", cnt)
	return arr
}

func TestSelectSort(t *testing.T) {
	arr := []int{9, 2, 5, 4, 1, 7, 6, 0, 8, 6}
	arr = SelectSort(arr)
	for i := range arr {
		fmt.Printf(" %v ", arr[i])
	}
}
