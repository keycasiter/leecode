package main

import (
	"fmt"
	"testing"
)

//循环
func TestInsertSort1(t *testing.T) {
	nums := []int{2, 3, 1, 6, 9, 8, 4, 5, 7}

	//遍历n次，每次都把后面未排序的数据和前面已排序好的元素做对比，并插入到前面已排序好的区间内
	//i代表每次数据规模开始的地方，它是递增的，i前面的是已排好序的，i后面的是未排好序的
	for i := 0; i < len(nums); i++ {
		//j代表每一次遍历要对比的元素
		//由于上面提到 i 前面都是排序完成的，i后面都是未排序的
		//因此 j = i 是选取未排序的第一个元素来进行和i 前面未排序的元素依次对比完成本次排序
		for j := i; j > 0; j-- {
			if nums[j] < nums[j-1] { // > 把较小的往前移; < 把较大的往前移
				//交换元素
				tmp := nums[j]
				nums[j] = nums[j-1]
				nums[j-1] = tmp
			}
		}
	}
	fmt.Println(nums)
}

//递归
func TestInsertSort2(t *testing.T) {
	nums := []int{2, 3, 1, 6, 9, 8, 4, 5, 7}
	insertSort(nums, 1)
	fmt.Println(nums)
}

func insertSort(nums []int, i int) {
	//当i超过nums长度就停止
	if i <= 0 || i >= len(nums) {
		return
	}

	//j代表每一次遍历要对比的元素
	//由于上面提到 i 前面都是排序完成的，i后面都是未排序的
	//因此 j = i 是选取未排序的第一个元素来进行和i 前面未排序的元素依次对比完成本次排序
	for j := i; j > 0; j-- {
		if nums[j] < nums[j-1] { // > 把较小的往前移; < 把较大的往前移
			//交换元素
			tmp := nums[j]
			nums[j] = nums[j-1]
			nums[j-1] = tmp
		}
	}
	//每一轮排序完，i递增，从下一个继续开始
	i++
	insertSort(nums, i)
}
