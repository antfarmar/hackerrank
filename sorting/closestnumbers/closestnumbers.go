package main

import (
	"bufio"
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

	// Parse input and populate map
	nums := []int{}
	for _, n := range cases {
		i, _ := strconv.Atoi(n)
		nums = append(nums, i)
	}
	sort.Ints(nums)
	fmt.Println(nums)

	// Find minimum difference
	min := int(^uint(0) >> 1) // maxint
	for i := 1; i < len(nums); i++ {
		diff := abs(nums[i] - nums[i-1])
		if diff < min {
			min = diff
		}
	}

	// Write the terms with min diff
	stdout := bufio.NewWriter(os.Stdout)
	for i := 1; i < len(nums); i++ {
		diff := abs(nums[i] - nums[i-1])
		if diff == min {
			stdout.WriteString(fmt.Sprint(nums[i-1], nums[i], " "))
		}
	}
	stdout.Flush()
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
