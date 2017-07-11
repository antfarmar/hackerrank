package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
)

var (
	//file := os.Stdin
	file, _     = os.Open(`.\data\loveletter.txt`)
	rawBytes, _ = ioutil.ReadAll(file)
	data        = bytes.Fields(rawBytes)[1:]
)

func main() {
	defer file.Close()

	for _, str := range data {
		fmt.Printf("%v\n", minOpsPalin(str))
	}
}

func minOpsPalin(s []byte) (cnt int) {
	i, j, mid := 0, len(s)-1, len(s)/2
	for ; i < mid; i, j = i+1, j-1 {
		if s[i] > s[j] {
			cnt += int(s[i] - s[j])
		} else if s[j] > s[i] {
			cnt += int(s[j] - s[i])
		}
	}
	return
}

/*
Sample Input
4
abc
abcba
abcd
cba

Sample Output
2
0
4
2
*/
