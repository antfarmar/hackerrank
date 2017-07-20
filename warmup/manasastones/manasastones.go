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
	file, _ := os.Open(`.\data\manasastones.txt`)
	rawBytes, _ := ioutil.ReadAll(file)

	cases := strings.Fields(string(rawBytes))[1:]

	for i := 0; i <= len(cases)-3; i += 3 {
		numStones, _ := strconv.Atoi(cases[i+0])
		diffA, _ := strconv.Atoi(cases[i+1])
		diffB, _ := strconv.Atoi(cases[i+2])
		lastStone(numStones, diffA, diffB)
	}
}

func lastStone(numStones, diffA, diffB int) {
	if numStones == 1 {
		fmt.Println("0")
		return
	}

	if diffA == diffB {
		fmt.Printf("%v\n", diffA*(numStones-1))
		return
	}

	if diffA > diffB {
		diffA, diffB = diffB, diffA
	}

	last := diffA * (numStones - 1)

	for i := 0; i < numStones; i++ {
		fmt.Printf("%v ", last-i*diffA+i*diffB)
	}

	fmt.Println()
}
