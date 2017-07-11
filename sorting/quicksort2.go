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

	numsSorted := qsort(nums)
	fmt.Println(toString(numsSorted))
}

func qsort(nums []int) (sorted []int) {
	pivot := nums[0]
	low, high := partition(pivot, nums)
	if len(low) > 1 {
		low = qsort(low)
		fmt.Println(toString(low))
	}
	if len(high) > 1 {
		high = qsort(high)
		fmt.Println(toString(high))
	}
	sorted = append(low, pivot)
	sorted = append(sorted, high...)
	return
}

func partition(pivot int, nums []int) (low, high []int) {
	for _, n := range nums[1:] {
		if n < pivot {
			low = append(low, n)
		} else {
			high = append(high, n)
		}
	}
	return
}

func toString(a []int) string {
	return fmt.Sprint(strings.Trim(fmt.Sprint(a), "[]"))
}
