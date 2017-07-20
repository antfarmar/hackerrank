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
	//file := os.Stdin
	file, _     = os.Open(`..\data.txt`)
	rawBytes, _ = ioutil.ReadAll(file)
	data        = parseInts(strings.Fields(string(rawBytes)))
)

func main() {
	defer file.Close()
	data = data[1:]
	size := 0

	for i := 0; i < len(data); i += size + 1 {
		size = data[i]
		nums := data[i+1 : i+size+1] // better to use a map
		fmt.Println(cntPairs(nums))
	}

}

func cntPairs(A []int) (numPairs uint64) {
	sort.Ints(A)
	sameNum := uint64(1)
	for i := 1; i < len(A); i++ {
		if A[i] == A[i-1] {
			sameNum++
		} else {
			numPairs += sameNum * (sameNum - 1)
			sameNum = 1
		}
	}
	numPairs += sameNum * (sameNum - 1)
	return numPairs
}

func parseInts(ss []string) (is []int) {
	is = make([]int, len(ss))
	for n, s := range ss {
		is[n], _ = strconv.Atoi(s)
	}
	return
}
