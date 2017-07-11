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
	file, _ := os.Open(`.\data\halloweenparty.txt`)
	rawBytes, _ := ioutil.ReadAll(file)

	cases := strings.Fields(string(rawBytes))[1:]

	for i := range cases {
		k, _ := strconv.Atoi(cases[i])
		switch k % 2 {
		case 0:
			fmt.Println(k * k / 4)
		case 1:
			fmt.Println(k / 2 * (k/2 + 1))
		}

	}
}
