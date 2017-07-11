package main

import (
	"fmt"
	"io/ioutil"
	"os"
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

	// Partition subslices until swapIndex == midIndex
	midIndex := len(nums) / 2
	first, last := 0, len(nums)-1
	swapIndex := partition(nums, first, last)

	for swapIndex != midIndex {
		if swapIndex > midIndex {
			last = swapIndex - 1
		} else {
			first = swapIndex + 1
		}
		swapIndex = partition(nums, first, last)
		//fmt.Println(nums, first, last)
	}

	fmt.Println(nums[midIndex])
}

func partition(a []int, first int, last int) int {
	swapIndex := first
	for i := first; i <= last; i++ {
		if a[i] < a[last] {
			swap(a, i, swapIndex)
			swapIndex++
		}
	}
	swap(a, last, swapIndex)
	return swapIndex
}

func swap(a []int, i, j int) {
	a[i], a[j] = a[j], a[i]
}
