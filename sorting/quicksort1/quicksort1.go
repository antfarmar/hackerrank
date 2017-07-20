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

	ansStr := fmt.Sprint(partition(nums[0], nums))
	fmt.Println(strings.Trim(ansStr, "[]"))
}

func partition(pivot int, nums []int) (part []int) {
	low, high := []int{}, []int{}
	for _, n := range nums {
		if n < pivot {
			low = append(low, n)
		} else {
			high = append(high, n)
		}
	}
	//part = append(low, pivot)
	part = append(low, high...)
	return
}
