package main

import (
	"fmt"
	"testing"
)

// 遍历
func TestSelectSort1(t *testing.T) {
	nums := []int{2, 3, 1, 6, 9, 8, 4, 5, 7}

	//这里的for循环是进行n次遍历，每一次的目的是把当前数据规模中最大或者最小的元素找到
	//然后将找到（选择）的这个元素放到当前数据范围的最前面，这样随着遍历的推移，每一次都会找到当前数据规模最大或者最小的元素
	//靠前面的数据已经排好序，靠后面的就是还没被排序的部分
	for i := 0; i < len(nums); i++ {
		//每次遍历取最开头的元素作为游标对象进行对比，通过它来找到最大或者最小的那个元素
		cmp := nums[i]
		//idx用来记录一次遍历后最大或最小元素的索引位置，为后面做交换使用的
		idx := i
		//这里从i+1开始就可以，因为cmp就是i位置的元素，不用和自己对比了
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < cmp { //对比元素大小, < 这里是找到最小的， > 的话是找到最大的
				cmp = nums[j]
				idx = j
			}
		}
		//交换元素
		//此时的cmp是这轮遍历后最大或者最小的那个元素值，idx是这个元素的索引，我们要把它放到本次数据范围的最前面
		tmp := nums[idx]
		nums[idx] = nums[i]
		nums[i] = tmp
	}
	fmt.Println(nums)
}

// 递归
func TestSelectSort2(t *testing.T) {
	nums := []int{2, 3, 1, 6, 9, 8, 4, 5, 7}

	selectSort(nums, 0)

	fmt.Println(nums)
}

func selectSort(nums []int, i int) {
	//i是数据规模开始的位置，i == len -1 代表当前只有一个元素，此时不需要排序了
	if i == len(nums)-1 {
		return
	}

	//每次遍历取最开头的元素作为游标对象进行对比，通过它来找到最大或者最小的那个元素
	cmp := nums[i]
	//idx用来记录一次遍历后最大或最小元素的索引位置，为后面做交换使用的
	idx := i
	//这里从i+1开始就可以，因为cmp就是i位置的元素，不用和自己对比了
	for j := i + 1; j < len(nums); j++ {
		if nums[j] < cmp { //对比元素大小, < 这里是找到最小的， > 的话是找到最大的
			cmp = nums[j]
			idx = j
		}
	}
	//交换元素
	//此时的cmp是这轮遍历后最大或者最小的那个元素值，idx是这个元素的索引，我们要把它放到本次数据范围的最前面
	tmp := nums[idx]
	nums[idx] = nums[i]
	nums[i] = tmp

	//每一次遍历都找到最大或者最小的元素并将其放在了当前数据范围的最开头
	//所以每次i前面都是排序好的，下一次排序从i的下一个开始
	i++
	selectSort(nums, i)
}

// 递归 + 二元选择、剪枝优化
func TestSelectSort3(t *testing.T) {
	nums := []int{2, 3, 1, 6, 9, 8, 4, 5, 7}

	selectSort3(nums, 0)

	fmt.Println(nums)
}

func selectSort3(nums []int, i int) {
	//i是数据规模开始的位置，i == len -1 代表当前只有一个元素，此时不需要排序了
	if i == len(nums)-1 {
		return
	}

	//每次遍历取最开头的元素作为游标对象进行对比，通过它来找到最大或者最小的那个元素
	min := nums[i]
	max := nums[i]
	//idx用来记录一次遍历后最大或最小元素的索引位置，为后面做交换使用的
	minIdx := i             //这里取本次数据范围的第一个为最小值，因为最小的会被放在第一个
	maxIdx := len(nums) - 1 //这里取最后一个为最大的对比，因为最大的会被放在最后一个
	//这里从i+1开始就可以，因为cmp就是i位置的元素，不用和自己对比了
	for j := i + 1; j < len(nums); j++ {
		if nums[j] < min { //对比元素大小, < 这里是找到最小的
			min = nums[j]
			minIdx = j
		}
		if nums[j] > max { //对比元素大小, > 这里是找到最大的
			max = nums[j]
			maxIdx = j
		}
	}

	//剪枝优化
	//最小位置的元素已经是最小，且最大位置的元素已经是最大，则说明剩余数据范围的元素都已经排序好了，可以停止遍历
	if minIdx == i && maxIdx == len(nums)-1 {
		return
	}

	//交换元素
	//此时的min、max是这轮遍历后最大或者最小的那个元素值
	//minIdx是这次遍历最小元素的索引，我妈要把它放到本次数据范围的最前面
	minTmp := nums[minIdx]
	nums[minIdx] = nums[i]
	nums[i] = minTmp
	//maxIdx是这次遍历最大元素的索引，我们要把它放到本次数据范围的最后面
	maxTmp := nums[maxIdx]
	nums[maxIdx] = nums[len(nums)-1]
	nums[len(nums)-1] = maxTmp
	//每一次遍历都找到最大或者最小的元素并将其放在了当前数据范围的最开头
	//所以每次i前面都是排序好的，下一次排序从i的下一个开始
	i++
	selectSort3(nums, i)
}
