package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

// MAIN ==================================================

func main() {
	//input,_ := ioutil.ReadAll(os.Stdin)
	input, _ := ioutil.ReadFile(`data\anagram.txt`)
	for _, words := range bytes.Fields(input)[1:] {
		fmt.Println(minEdit(words))
	}
}

// SOLUTION ==============================================

func minEdits(wordPair []byte) (edits int) {
	if len(wordPair)&1 == 1 {
		return -1
	}
	mid := len(wordPair) / 2

Outer:
	for _, x := range wordPair[:mid] {
		for j, y := range wordPair[mid:] {
			if y == x {
				wordPair[mid+j] = '.'
				continue Outer
			}
		}
		edits++
	}
	return
}
