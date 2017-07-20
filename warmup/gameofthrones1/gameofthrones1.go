package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	//file := os.Stdin
	file, _ := os.Open(`.\data\gameofthrones1.txt`)
	in, _ := ioutil.ReadAll(file)
	key := strings.TrimSpace(string(in))

	ans := "YES"
	chrCnts := make(map[rune]int)
	oddCnt := 0

	for _, chr := range key {
		chrCnts[chr]++
	}

	fmt.Fprintln(os.Stderr, chrCnts)

	if len(chrCnts) > 1 {
		for _, cnt := range chrCnts {
			if cnt%2 == 1 {
				oddCnt++
			}
			if oddCnt > 1 {
				ans = "NO"
				break
			}
		}
	}

	fmt.Println(ans)
}
