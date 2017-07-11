package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
)

// MAIN ==================================================

func main() {
	//input,_ := ioutil.ReadAll(os.Stdin)
	input, _ := ioutil.ReadFile(`data\biggerisgreater.txt`)
	for _, word := range bytes.Fields(input)[1:] {
		fmt.Fprintf(os.Stdout, "%s\n", lexicographicallyNext(word))
	}
}

// SOLUTION ==============================================

func lexicographicallyNext(word []uint8) (next []uint8) {
	next = []uint8("no answer")

	for i := len(word) - 2; i >= 0; i-- {
		for j := len(word) - 1; j > i; j-- {
			if word[i] < word[j] {
				swap(word, i, j)
				reverse(word[i+1 : len(word)])
				next = word
				return
			}
		}
	}
	return
}

func swap(a []uint8, i, j int) {
	a[j], a[i] = a[i], a[j]
}

func reverse(a []uint8) {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}
