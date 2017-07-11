package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"strconv"
)

// MAIN ==================================================

func main() {
	//input, _ := ioutil.ReadAll(os.Stdin)
	input, _ := ioutil.ReadFile(`data\anagram.txt`)
	var buffer bytes.Buffer
	for _, words := range bytes.Fields(input)[1:] {
		buffer.WriteString(strconv.Itoa(minEdit(words)) + "\n")
	}
	io.WriteString(os.Stdout, buffer.String())
}

// SOLUTION ==============================================

func minEdits(wordPair []byte) (minEdits int) {
	if len(wordPair)&1 == 1 {
		return -1
	}

	mid := len(wordPair) / 2
	ltrCnt := make(map[byte]int, mid)

	for _, b := range wordPair[:mid] {
		ltrCnt[b]++
	}
	for _, b := range wordPair[mid:] {
		ltrCnt[b]--
	}

	for _, v := range ltrCnt {
		if v > 0 {
			minEdits += v
		}
	}
	return
}
