package utils

import "sort"

var check SliceInterface = &SortExistence{}

type SliceInterface interface {
	SliceIsExit(arr []string, one string) (index int, isExit bool)
	SliceRemoveInt(int, []int) []int
}
type SortExistence struct {
}

// SliceIsExit 判断切片中是否存在某一元素 string类型
func (s *SortExistence) SliceIsExit(arr []string, one string) (index int, isExit bool) {
	index = sort.SearchStrings(arr, one) //采用了二分查找，找到返回索引位置，没有找到返回可以插入的位置
	isExit = index != len(arr) && arr[index] == one
	return index, isExit
}

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
