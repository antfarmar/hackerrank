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

	qsort(sort.IntSlice(nums))
}

func qsort(a sort.Interface) {
	quicksort(a, 0, a.Len()-1)
}

func quicksort(a sort.Interface, first int, last int) {
	if first >= last {
		return
	}
	pivotIndex := partition(a, first, last, last)
	fmt.Println(strings.Trim(fmt.Sprint(a), "[]"))
	quicksort(a, first, pivotIndex-1)
	quicksort(a, pivotIndex+1, last)
}

func partition(a sort.Interface, first int, last int, pivotIndex int) int {
	swapIndex := first
	for i := first; i <= last; i++ {
		if a.Less(i, pivotIndex) {
			a.Swap(i, swapIndex)
			swapIndex++
		}
	}
	a.Swap(pivotIndex, swapIndex)
	return swapIndex
}
