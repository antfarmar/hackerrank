package main

import (
	"bufio"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	rawBytes, _ := ioutil.ReadAll(os.Stdin)
	cases := strings.Fields(string(rawBytes))[1:]

	// Parse input and populate map
	cntMap := make(map[int]int)
	for i := 0; i < len(cases); i += 2 {
		i, _ := strconv.Atoi(cases[i])
		cntMap[i]++
	}

	// Write sum of counts to stdout
	sum := 0
	stdout := bufio.NewWriter(os.Stdout)
	for i := 0; i < 100; i++ {
		sum += cntMap[i]
		stdout.WriteString(strconv.Itoa(sum) + " ")
	}
	stdout.Flush()
}
