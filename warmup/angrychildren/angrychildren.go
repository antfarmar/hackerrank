package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	//file := os.Stdin
	file, _ := os.Open(`.\data\angrychildren.txt`)
	in := bufio.NewScanner(file)
	nums := []int{}

	//in.Scan()
	for in.Scan() {
		line := in.Text()
		i, _ := strconv.Atoi(line)
		nums = append(nums, i)
	}

	n := nums[0]
	k := nums[1]
	nums = nums[2:]
	sort.Ints(nums)

	fairestSet := minimizeUnfairness(nums, k)
	min, max := minAndMax(fairestSet)

	fmt.Fprintln(os.Stderr, n, k, fairestSet, max, min, max-min)
	fmt.Println(max - min)
}

func minimizeUnfairness(arr []int, k int) (ans []int) {
	beg, end, best := 0, k, 2<<30 // 2147483647 > 10^9
	for i := 0; i <= len(arr)-k; i++ {
		min, max := minAndMax(arr[i : i+k])
		if max-min <= best {
			best = max - min
			beg, end = i, i+k
		}
	}
	return arr[beg:end]
}

func minAndMax(a []int) (min, max int) {
	min, max = 2<<30, 0
	for _, n := range a {
		if n > max {
			max = n
		}
		if n < min {
			min = n
		}
	}
	return
}
