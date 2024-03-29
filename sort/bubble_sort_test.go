package main

import (
	"fmt"
	"testing"
)

//双循环
func TestBubbleSort1(t *testing.T) {
	nums := []int{2, 3, 1, 6, 9, 8, 4, 5, 7}

	//这里for循环表示要遍历的趟数
	//第一次遍历全部的[0,n]个，然后排序好的最后一个数是符合预期的（最大或者最小）
	//第二次遍历全部的[0,n-1]个，然后最后两个就被排序好了
	//以此类推
	//i代表每一趟参与遍历的元素数量
	for i := len(nums); i > 0; i-- {
		//这里for循环代表每一趟比较、交换的过程
		//每次都是从[0,n]的范围，只不过n会每次由外循环传递进来，并且每一趟最后一个会排序好，所以n的数量会递减
		//j代表每次遍历的下标，j<i-1 代表只遍历到数据规模的倒数第二个，因为要进行对比，这里目的是防止索引越界
		for j := 0; j < i-1; j++ {
			//nums[j]代表当前遍历到的元素，nums[j+1]代表下一个
			if nums[j+1] < nums[j] { //判断元素大小
				//交换过程
				tmp := nums[j+1]
				nums[j+1] = nums[j]
				nums[j] = tmp
			}
		}
	}
	//打印
	fmt.Println(nums)
}

//递归
func TestBubbleSort2(t *testing.T) {
	nums := []int{2, 3, 1, 6, 9, 8, 4, 5, 7}
	bubbleSort(nums, len(nums))
	//打印
	fmt.Println(nums)
}

func bubbleSort(nums []int, size int) {
	//不满足两个不需要比较
	if size < 2 {
		return
	}
	//这里for循环代表每一趟比较、交换的过程
	//每次都是从[0,n]的范围，只不过n会每次由外循环传递进来，并且每一趟最后一个会排序好，所以n的数量会递减
	//j代表每次遍历的下标，j<i-1 代表只遍历到数据规模的倒数第二个，因为要进行对比，这里目的是防止索引越界
	for j := 0; j < size-1; j++ {
		//nums[j]代表当前遍历到的元素，nums[j+1]代表下一个
		if nums[j+1] < nums[j] { //判断元素大小
			//交换过程
			tmp := nums[j+1]
			nums[j+1] = nums[j]
			nums[j] = tmp
		}
	}
	//每次数据规模减少1，因为每一次最靠后的都排序好了
	size--
	bubbleSort(nums, size)
}

//递归
func TestBubbleSort3(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	bubbleSort3(nums, len(nums))
	//打印
	fmt.Println(nums)
}

//优化
func bubbleSort3(nums []int, size int) {
	//不满足两个不需要比较
	if size < 2 {
		return
	}
	//增加优化标识，判断是否都排序过了
	sorted := true

	//这里for循环代表每一趟比较、交换的过程
	//每次都是从[0,n]的范围，只不过n会每次由外循环传递进来，并且每一趟最后一个会排序好，所以n的数量会递减
	//j代表每次遍历的下标，j<i-1 代表只遍历到数据规模的倒数第二个，因为要进行对比，这里目的是防止索引越界
	for j := 0; j < size-1; j++ {
		//nums[j]代表当前遍历到的元素，nums[j+1]代表下一个
		if nums[j+1] < nums[j] { //判断元素大小
			//交换过程
			tmp := nums[j+1]
			nums[j+1] = nums[j]
			nums[j] = tmp
			sorted = false
		}
	}
	//已经排序过的就跳过了
	if sorted {
		return
	}
	//每次数据规模减少1，因为每一次最靠后的都排序好了
	size--
	bubbleSort3(nums, size)
}
