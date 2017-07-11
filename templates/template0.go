package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	data := readData(1)
	fmt.Printf("%s\n", data)
}

func readData(in int) []string {
	b, err := []byte{}, error(nil)
	if in == 0 {
		b, err = ioutil.ReadAll(os.Stdin)
	} else {
		b, err = ioutil.ReadFile(`..\data.txt`)
	}
	if err != nil {
		panic(err)
	}
	return strings.Fields(string(b))
}
