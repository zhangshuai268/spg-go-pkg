package iarray

import (
	"reflect"
)

// IsContain 检查int是否在[]int数组中
func IsContain(items []int, item int) (bool, int) {
	for index, eachItem := range items {
		if eachItem == item {
			return true, index
		}
	}
	return false, -1
}

// Reverse 反转数组
func Reverse(s interface{}) {
	size := reflect.ValueOf(s).Len()
	swap := reflect.Swapper(s)
	for i, j := 0, size-1; i < j; i, j = i+1, j-1 {
		swap(i, j)
	}
}

// RemoveReplica []int数组去重
func RemoveReplica(slc []int) []int {
	result := make([]int, 0)
	tempMap := make(map[int]bool, len(slc))
	for _, e := range slc {
		if tempMap[e] == false {
			tempMap[e] = true
			result = append(result, e)
		}
	}
	return result
}

// Union 求并集
func Union(slice1, slice2 []int) []int {
	m := make(map[int]int)
	for _, v := range slice1 {
		m[v]++
	}

	for _, v := range slice2 {
		times, _ := m[v]
		if times == 0 {
			slice1 = append(slice1, v)
		}
	}
	return slice1
}

// Intersect 求交集
func Intersect(slice1, slice2 []int) []int {
	m := make(map[int]int)
	nn := make([]int, 0)
	for _, v := range slice1 {
		m[v]++
	}

	for _, v := range slice2 {
		times, _ := m[v]
		if times > 0 {
			nn = append(nn, v)
		}
	}
	return nn
}
