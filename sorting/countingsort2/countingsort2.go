package main

import (
	"bufio"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	file := os.Stdin
	//file, _ := os.Open(`..\data.txt`)
	rawBytes, _ := ioutil.ReadAll(file)
	cases := strings.Fields(string(rawBytes))[1:]

	// Parse input and populate map
	nums := []int{}
	cntMap := make(map[int]int)
	for _, n := range cases {
		i, _ := strconv.Atoi(n)
		nums = append(nums, i)
		cntMap[i]++
	}

	// Write map in inc. order to stdout
	stdout := bufio.NewWriter(os.Stdout)
	for i := 0; i < 100; i++ {
		cnt := cntMap[i]
		if cnt > 0 {
			str := strings.Repeat(strconv.Itoa(i)+" ", cnt)
			stdout.WriteString(str)
		}
	}
	stdout.Flush()
}
