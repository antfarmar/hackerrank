package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	//file := os.Stdin
	file, _ := os.Open(`.\data\sherlockbeast.txt`)
	rawBytes, _ := ioutil.ReadAll(file)
	cases := strings.Fields(string(rawBytes))[1:]

	for i := 0; i < len(cases); i++ {
		n, _ := strconv.Atoi(cases[i])
		fmt.Println(n, decentNumber(n))
	}
}

func decentNumber(n int) (ans string) {
	repeat3, repeat5 := 0, 0

	// Look first for x,y sth. n = 3x + 5y with x>0
	for i := 0; i < n; i++ {
		if (n-i)%3 == 0 && i%5 == 0 {
			repeat3 = i
			repeat5 = n - i
			break
		}
	}

	// If no match, check for n = 5x or no such x
	if repeat3 == 0 && repeat5 == 0 {
		if n%5 == 0 {
			repeat3 = n
		} else {
			return "-1"
		}
	}

	// Construct the decent number since it exists here
	var buf bytes.Buffer
	buf.Grow(repeat3 + repeat5)
	for i := 0; i < repeat5; i++ {
		buf.WriteString("5")
	}
	for i := 0; i < repeat3; i++ {
		buf.WriteString("3")
	}

	return buf.String()
}
