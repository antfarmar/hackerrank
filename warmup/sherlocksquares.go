package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	//file := os.Stdin
	file, _ := os.Open(`.\data\sherlocksquares.txt`)
	rawBytes, _ := ioutil.ReadAll(file)

	cases := strings.Fields(string(rawBytes))[1:]

	for i := 0; i <= len(cases)-2; i += 2 {
		a, _ := strconv.ParseFloat(cases[i+0], 64)
		b, _ := strconv.ParseFloat(cases[i+1], 64)
		fmt.Println(numSquares(a, b))
	}
}

func numSquares(min, max float64) (cnt int) {
	start := math.Ceil(math.Sqrt(min))
	for ; start*start <= max; start++ {
		cnt++
	}
	return
}
