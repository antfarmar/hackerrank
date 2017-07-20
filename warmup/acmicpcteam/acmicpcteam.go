package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	//file := os.Stdin
	file, _ := os.Open(`.\data\acmicpcteam.txt`)
	data, _ := ioutil.ReadAll(file)
	bitMaps := strings.Fields(string(data))[2:]
	n, m := len(bitMaps), len(bitMaps[0])

	// map{#topics a team knows : #teams that know this many topics}
	topicsTeamCnt := make(map[int]int)

	for i := range bitMaps {
		for j := i + 1; j < len(bitMaps); j++ {
			numTopics := orCnt(bitMaps[i], bitMaps[j])
			topicsTeamCnt[numTopics]++
		}
	}

	max, val := 0, 0
	for k, v := range topicsTeamCnt {
		if k > max {
			max = k
			val = v
		}
	}

	fmt.Fprintf(os.Stderr, "%v %v %v\n%v\n", n, m, bitMaps, topicsTeamCnt)
	fmt.Printf("%v\n%v\n", max, val)
}

func orCnt(bs1, bs2 string) (cnt int) {
	for i := range bs1 {
		if bs1[i] == '1' || bs2[i] == '1' {
			cnt++
		}
	}
	return
}

func cntOnes(bs string) (cnt int) {
	for _, b := range bs {
		if b == '1' {
			cnt++
		}
	}
	return
}
