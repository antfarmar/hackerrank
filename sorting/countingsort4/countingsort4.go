package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Data struct {
	num int
	str string
}

func main() {
	//file := os.Stdin
	file, _ := os.Open(`..\data.txt`)
	rawBytes, _ := ioutil.ReadAll(file)
	cases := strings.Fields(string(rawBytes))[1:]

	// Parse input and populate map
	var cnt [100]int
	var data = make([]Data, len(cases)/2)
	for i := 0; i < len(cases); i += 2 {
		n, _ := strconv.Atoi(cases[i])
		cnt[n]++
		data[i/2] = Data{n, cases[i+1]}
	}

	//fmt.Printf("%v\n%v\n", cnt, data)

	// Make indexing helper array
	var idx [101]int //idx := make([]int, len(cntMap)+1)
	for i := 1; i <= len(cnt); i++ {
		idx[i] = cnt[i-1] + idx[i-1]
	}

	//fmt.Println(idx)

	// Write sorted strings to stdout
	pos := 0
	sorted := make([]string, idx[len(idx)-1])
	for i := 0; i < 100; i++ {
		if cnt[i] > 0 {
			for j := 0; j < len(data); j++ {
				if i == data[j].num {
					if j > len(data)/2 {
						sorted[idx[pos]] = data[j].str
					} else {
						sorted[idx[pos]] = "-"
					}
					idx[pos]++
				}
			}
			pos++
		}
	}
	fmt.Println(strings.Join(sorted, " "))
}
