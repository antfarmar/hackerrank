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

/*
 XOR Property:
 Even appearances of #n == 0
 Odd  appearances of #n == n

In only odd sized arrays do we get #'s to appear
an odd # of times in the set of all contiguous subarrays.
They are at located at even indexes only.
*/
func xorAllSubarrays(A []int) (xor int) {
	if len(A)&1 == 1 {
		for i := 0; i < len(A); i += 2 {
			xor ^= A[i]
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
