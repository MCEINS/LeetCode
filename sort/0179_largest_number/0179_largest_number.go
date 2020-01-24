package main

import (
	"fmt"
	"sort"
	"strconv"
)

func largestNumber(nums []int) string {
	size := len(nums)
	b := make(bytes, size)
	resSize := 0

	for i := range b {
		b[i] = []byte(strconv.Itoa(nums[i]))
		resSize += len(b[i])
	}

	sort.Sort(b)

	res := make([]byte, 0, resSize)
	for i := size - 1; i >= 0; i-- {
		res = append(res, b[i]...)
	}

	i := 0
	for i < resSize-1 && res[i] == '0' {
		i++
	}

	return string(res[i:])

}

type bytes [][]byte

func (b bytes) Len() int {
	return len(b)
}

func (b bytes) Less(i, j int) bool {
	size := len(b[i]) + len(b[j])

	bij := make([]byte, 0, size)
	bij = append(bij, b[i]...)
	bij = append(bij, b[j]...)

	bji := make([]byte, 0, size)
	bji = append(bji, b[j]...)
	bji = append(bji, b[i]...)

	for k := 0; k < size; k++ {
		if bij[k] < bji[k] {
			return true
		} else if bij[k] > bji[k] {
			return false
		}
	}
	return false
}

func (b bytes) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func main() {
	nums := []int{1, 2, 3}
	res := largestNumber(nums)
	fmt.Println(res)
}
