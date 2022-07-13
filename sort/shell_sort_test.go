package main

import "testing"

func TestShellSort1(t *testing.T) {
	nums := []int{2, 3, 1, 6, 9, 8, 4, 5, 7}
	size := len(nums)

	//每次减为原来数据规模的一般，取整，最后一定为1
	//groupSize为拆分的组数量
	for groupSize := size / 2; groupSize > 0; groupSize = groupSize / 2 {
		//每个组分别执行插入排序操作
		for i := 0; i < groupSize; i++ {

		}
	}
}

func groupSort(nums []int, size, startIdx, groupSize int) {

}
