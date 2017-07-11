package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	//data := readData("stdin")
	data := readData(`.\data\utopiantree.txt`)

	for _, s := range data[1:] {
		n, _ := strconv.Atoi(s)
		fmt.Printf("%v\n", growth(n))
	}
}

func growth(cycles int) (height int) {
	height = 1
	for n := 1; n <= cycles; n++ {
		switch n % 2 {
		case 0:
			height += 1
		case 1:
			height *= 2
		}
	}
	return height
}

func readData(name string) []string {
	b, err := []byte{}, error(nil)
	if name == "stdin" {
		b, err = ioutil.ReadAll(os.Stdin)
	} else {
		b, err = ioutil.ReadFile(name)
	}
	if err != nil {
		panic(err)
	}
	return strings.Fields(string(b))
}
