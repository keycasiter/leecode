package sort

import (
	"fmt"
	"testing"
)

func BubbleSort(arr []int) []int {
	cnt := 0
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				temp := arr[j+1]
				arr[j+1] = arr[j]
				arr[j] = temp
			}
			cnt++
		}
	}
	fmt.Println("cnt -> ", cnt)
	return arr
}

func BubbleSortV2(arr []int) []int {
	cnt := 0
	n := len(arr) - 1
	for i := 0; i < len(arr); i++ {
		for j := 0; j < n; j++ {
			if arr[j] > arr[j+1] {
				temp := arr[j+1]
				arr[j+1] = arr[j]
				arr[j] = temp
			}
			cnt++
		}
		n--
	}
	fmt.Println("cnt -> ", cnt)
	return arr
}

func BubbleSortV3(arr []int) []int {
	cnt := 0
	n := len(arr)
	flag := true
	for i := 0; i < n && flag; i++ {
		flag = false
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				flag = true
			}
			cnt++
		}
	}
	fmt.Println(cnt)
	return arr
}

func TestBubbleSort(t *testing.T) {
	arr := []int{2, 6, 7, 1, 4, 3, 8, 9, 0, 5}
	arr = BubbleSortV3(arr)
	for i := range arr {
		fmt.Printf(" %v ", arr[i])
	}
}
