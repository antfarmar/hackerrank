package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
)

var (
	//fileIn  = os.Stdin
	fileOut   = os.Stdout
	fileIn, _ = os.Open(`data\anagram.txt`)
	rw        = bufio.NewReadWriter(bufio.NewReader(fileIn), bufio.NewWriter(fileOut))
)

// MAIN ==================================================

func main() {
	defer fileOut.Close()
	defer fileIn.Close()
	defer rw.Flush()

	var input []byte
	for line, err := rw.ReadSlice('\n'); err != io.EOF; {
		line, err = rw.ReadSlice('\n')
		input = append(input, line...)
	}

	words := bytes.Fields(input)
	for _, word := range words {
		rw.WriteString(fmt.Sprintln(minEdit(word)))
	}
}

// SOLUTION ==============================================

func minEdit(wordPair []byte) (minEdits int) {
	if len(wordPair)&1 == 1 {
		return -1
	}

	ltrCnt := make(map[byte]int, len(wordPair))

	// Map word 2 to unique keys (uppercase)
	for i, b := range wordPair {
		if i < len(wordPair)/2 {
			ltrCnt[b]++
		} else {
			ltrCnt[b-32]++
		}
	}

	for i := byte('a'); i <= 'z'; i++ {
		minEdits += abs(ltrCnt[i] - ltrCnt[i-32])
	}
	minEdits /= 2
	return
}

func abs(n int) int {
	if n < 0 {
		n = -n
	}
	return n
}
