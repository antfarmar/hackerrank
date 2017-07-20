package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

var (
	//file := os.Stdin
	file, _     = os.Open(`..\data.txt`)
	rawBytes, _ = ioutil.ReadAll(file)
	data        = parseInts(strings.Fields(string(rawBytes)))
)

func main() {
	defer file.Close()

	size := data[0]
	rots := data[1] % size
	qrys := data[size+3:]
	nums := data[3 : size+3]

	for _, q := range qrys {
		index := (q - rots + size) % size
		fmt.Println(nums[index])
	}

	fmt.Fprintln(os.Stderr, size, rots, nums, qrys)
}

func parseInts(ss []string) (is []int) {
	is = make([]int, len(ss))
	for n, s := range ss {
		is[n], _ = strconv.Atoi(s)
	}
	return
}
