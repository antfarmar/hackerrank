package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	//data := readData("stdin")
	data := readData(`..\data.txt`)
	nums := parseInts(data)

	V, n, ar := nums[0], nums[1], nums[2:]
	fmt.Printf("%v\n", binSearch(ar, 0, n, V))
}

func binSearch(nums []int, lo, hi, key int) int {
	if lo == hi {
		return -1
	}
	mid := (lo + hi) / 2
	val := nums[mid]
	if key < val {
		return binSearch(nums, lo, mid, key)
	} else if key > val {
		return binSearch(nums, mid+1, hi, key)
	} else { // key == val
		return mid
	}
}

func parseInts(ss []string) (is []int) {
	is = make([]int, len(ss))
	for n, s := range ss {
		is[n], _ = strconv.Atoi(s)
	}
	return
}

func readData(name string) []string {
	b, err := []byte{}, error(nil)
	if name == "stdin" {
		b, err = ioutil.ReadAll(os.Stdin)
	} else {
		b, err = ioutil.ReadFile(name)
	}
	if err != nil {
		panic(err)
	}
	return strings.Fields(string(b))
}
