package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

var (
	//file        = os.Stdin
	file, _     = os.Open(`.\data\andproduct.txt`)
	rawBytes, _ = ioutil.ReadAll(file)
	data        = strings.Fields(string(rawBytes))[1:]
)

func main() {
	// For each test case pair...
	for i := 0; i < len(data); i += 2 {
		lo, _ := strconv.ParseUint(data[i+0], 10, 64)
		hi, _ := strconv.ParseUint(data[i+1], 10, 64)
		fmt.Println(andProduct(lo, hi))
	}
}

func andProduct(lo, hi uint64) (prod uint64) {
	printBits(lo, hi)
	prod = lo & hi
	pows2 := uint64(1)
	lo += pows2
	for lo < hi {
		prod &= lo
		pows2 <<= 1
		lo += pows2
	}
	return
}

func printBits(lo, hi uint64) {
	for i := lo; i <= hi; i++ {
		fmt.Printf("%b %v\n", i, i)
	}
}
