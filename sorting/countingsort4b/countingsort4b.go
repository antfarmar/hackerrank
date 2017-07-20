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
	file, _ := os.Open(`..\data.txt`)
	rawBytes, _ := ioutil.ReadAll(file)
	cases := strings.Fields(string(rawBytes))[1:]

	// Parse input and populate map
	cntMap := make(map[int]int)
	datMap := make(map[int][]string)
	for i := 0; i < len(cases); i += 2 {
		n, _ := strconv.Atoi(cases[i])
		cntMap[n]++
		datMap[n] = append(datMap[n], cases[i+1])
	}

	fmt.Println(datMap)

	// Make indexing helper array
	idx := make([]int, len(cntMap)+1)
	for i := 1; i <= len(cntMap); i++ {
		idx[i] = cntMap[i-1] + idx[i-1]
	}
	fmt.Println(idx)

	// Write sum of counts to stdout
	//stdout := bufio.NewWriter(os.Stdout)
	//pos := 0
	var sorted [100]string
	for i := 0; i <= 100; i++ {
		if cntMap[i] > 0 {
			sorted[i] = strings.Join(datMap[i], " ") //append(sorted, datMap[i]...)
		}
	}
	fmt.Println(sorted)
	//stdout.Flush()
}
