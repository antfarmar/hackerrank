package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

var (
	//file        = os.Stdin
	file, _     = os.Open(`.\data\countergame.txt`)
	rawBytes, _ = ioutil.ReadAll(file)
	data        = strings.Fields(string(rawBytes))[1:]
)

func main() {
	// For each test case...
	for _, d := range data {
		n, _ := strconv.ParseUint(d, 10, 64)
		moveCnt := countGame(n)
		fmt.Println([]string{"Richard", "Louise"}[moveCnt%2])
	}
}

func countGame(n uint64) (cnt int) {
	for n != 1 {
		if isPower2(n) {
			println("Y", n)
			n = n >> 1
		} else {
			println("N", n)
			log2 := uint64(math.Log2(float64(n)))
			n <<= (64 - log2)
			n >>= (64 - log2)

			//println(n, log2)
			//log2 := math.Log2(float64(n))
			//pow2 := uint64(math.Pow(2, math.Floor(log2)))
			//println(n&^pow2)
			//n = n&^pow2 // andNot == bitclear
		}
		cnt++
	}
	return
}

func isPower2(n uint64) (ans bool) {
	for i := uint(0); i < 64; i++ {
		if n == 1<<i {
			ans = true
			break
		}
	}
	return
}
