package main

import (
	"bufio"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	//file := os.Stdin
	file, _ := os.Open(`..\data.txt`)
	rawBytes, _ := ioutil.ReadAll(file)
	cases := strings.Fields(string(rawBytes))[1:]

	// Parse input and populate map
	cntMap := make(map[int]int)
	for _, n := range cases {
		i, _ := strconv.Atoi(n)
		cntMap[i]++
	}

	// Write counts to stdout
	stdout := bufio.NewWriter(os.Stdout)
	for i := 0; i < 100; i++ {
		stdout.WriteString(strconv.Itoa(cntMap[i]) + " ")
	}
	stdout.Flush()
}
