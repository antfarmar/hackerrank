package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	//file := os.Stdin
	file, _ := os.Open(`.\data\chocolatefeast.txt`)
	rawBytes, _ := ioutil.ReadAll(file)

	cases := parseInts(strings.Fields(string(rawBytes))[1:])

	for i := 0; i <= len(cases)-3; i += 3 {
		NCM := cases[i : i+3]
		fmt.Println(numChocs(NCM[0], NCM[1], NCM[2]))
	}
}

func numChocs(cash, cost, disc int) (total int) {
	total = cash / cost
	wraps := total
	for wraps >= disc {
		total += wraps / disc
		wraps = wraps/disc + wraps%disc
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
