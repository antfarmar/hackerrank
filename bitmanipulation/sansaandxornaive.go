package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

var (
	//file = os.Stdin
	file, _     = os.Open(`.\data\sansaandxor.txt`)
	rawBytes, _ = ioutil.ReadAll(file)
	data        = parseInts(strings.Fields(string(rawBytes)))[1:]
)

func main() {
	// For each test case...
	for i := 0; i < len(data); i += data[i] + 1 {
		array := data[i+1 : data[i]+i+1]
		fmt.Println(xorAllSubarrays(array))
	}
}

func xorAllSubarrays(A []int) (xor int) {
	for n := 1; n <= len(A); n++ {
		for i := 0; i <= len(A)-n; i++ {
			xor ^= xorArray(A[i : i+n])
			//fmt.Println(A[i : i+n])
		}
	}
	return
}

func xorArray(A []int) (xor int) {
	for i := 0; i < len(A); i++ {
		xor ^= A[i]
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
