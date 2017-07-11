package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
)

var (
	//file := os.Stdin
	file, _     = os.Open(`.\data\alternatingchars.txt`)
	rawBytes, _ = ioutil.ReadAll(file)
	data        = bytes.Fields(rawBytes)[1:]
)

func main() {
	defer file.Close()

	for _, str := range data {
		fmt.Printf("%v\n", minDelCnt(str))
	}
}

// Count deletions
func minDelCnt(s []byte) (cnt int) {
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			cnt++
		}
	}
	return
}

// Construct the alternating string
func minDelStr(s []byte) (str string) {
	b := make([]byte, 1)

	b[0] = s[0]
	for i := 1; i < len(s); i++ {
		if s[i] != b[len(b)-1] {
			b = append(b, s[i])
		}
	}
	return string(b)
}

/*
Sample Input
5
AAAA
BBBBB
ABABABAB
BABABA
AAABBB

Sample Output
3
4
0
0
4
*/
