package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	//file := os.Stdin
	file, _ := os.Open(`..\data.txt`)
	rawBytes, _ := ioutil.ReadAll(file)
	cases := strings.Fields(string(rawBytes))[1:]

	nums := []int{}
	for _, n := range cases {
		i, _ := strconv.Atoi(n)
		nums = append(nums, i)
	}
	nums2 := make([]int, len(nums))
	copy(nums2, nums)

	shiftCnt := insertionSort(nums)
	swapCnt := qsort(sort.IntSlice(nums2))

	fmt.Println(shiftCnt - swapCnt)
}

func qsort(a sort.Interface) int {
	return quicksort(a, 0, a.Len()-1)
}

func quicksort(a sort.Interface, first, last int) int {
	if first >= last {
		return 0
	}
	pivotIndex, swapCnt := partition(a, first, last, last)
	cntL := quicksort(a, first, pivotIndex-1)
	cntR := quicksort(a, pivotIndex+1, last)
	return swapCnt + cntL + cntR

}

func partition(a sort.Interface, first int, last int, pivotIndex int) (int, int) {
	swapIndex, swapCnt := first, 0
	for i := first; i <= last; i++ {
		if a.Less(i, pivotIndex) {
			a.Swap(i, swapIndex)
			swapIndex++
			swapCnt++
		}
	}
	a.Swap(pivotIndex, swapIndex)
	swapCnt++
	return swapIndex, swapCnt
}

func insertionSort(array []int) (shiftCnt int) {
	for i := 1; i < len(array); i++ {
		for j := i; j > 0 && array[j] < array[j-1]; j-- {
			array[j], array[j-1] = array[j-1], array[j]
			shiftCnt++
		}
	}
	return
}
