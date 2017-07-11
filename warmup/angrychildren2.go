package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

var (
	//file        = os.Stdin
	file, _     = os.Open(`.\data\angrychildren.txt`)
	rawBytes, _ = ioutil.ReadAll(file)
	nums        = parseInts(strings.Fields(string(rawBytes)))
)

func main() {
	defer file.Close()

	n := nums[0]
	k := nums[1]
	nums = nums[2:]
	sort.Ints(nums)

	fmt.Println(minimizeUnfairness(nums, n, k))
	fmt.Fprintln(os.Stderr, n, k, nums)
}

func minimizeUnfairness(arr []int, n, k int) (min int) {
	min = 1 << 32
	for i := 0; i < n-k; i++ {
		diff := nums[i+k-1] - nums[i]
		if diff < min {
			min = diff
		}
	}
	return
}

func parseInts(ss []string) (is []int) {
	is = make([]int, len(ss))
	for n, s := range ss {
		is[n], _ = strconv.Atoi(s)
	}
	return
}
