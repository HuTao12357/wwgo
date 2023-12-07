package utils

import "sort"

var check SliceInterface = &SortExistence{}

type SliceInterface interface {
	SliceIsExit(arr []string, one string) (index int, isExit bool)
	SliceRemoveInt(int, []int) []int
	SliceDuplicates(arr []int) []int
	SliceBathRemoves(arr []int, index []int) []int
}
type SortExistence struct {
}

// SliceBathRemoves 切片批量删除(索引)
func (s *SortExistence) SliceBathRemoves(arr []int, index []int) []int {
	if len(arr) == 0 || len(index) == 0 {
		return arr
	}
	//索引逆序排序
	sort.Sort(sort.Reverse(sort.IntSlice(index)))
	//获取索引值
	k := index[0]
	arr = append(arr[:k], arr[k+1:]...)
	return s.SliceBathRemoves(arr, index[1:])
}

// SliceDuplicates 切片去重
func (s *SortExistence) SliceDuplicates(arr []int) []int {
	noMap := make(map[int]bool)
	res := make([]int, 0)
	for _, v := range arr {
		if noMap[v] == false {
			noMap[v] = true
			res = append(res, v)
		}
	}
	return res
}

// SliceIsExit 判断切片中是否存在某一元素 string类型
func (s *SortExistence) SliceIsExit(arr []string, one string) (index int, isExit bool) {
	index = sort.SearchStrings(arr, one) //采用了二分查找，找到返回索引位置，没有找到返回可以插入的位置
	isExit = index != len(arr) && arr[index] == one
	return index, isExit
}

// SliceRemoveInt remove
func (s *SortExistence) SliceRemoveInt(a int, arr []int) []int {
	var index int
	for k, v := range arr {
		if v == a {
			index = k
		}
	}
	arr = append(arr[:index], arr[index+1:]...) //...表示将切片或数组拆开，作为可变参数传递给函数
	return arr
}
