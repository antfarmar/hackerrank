package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	//file := os.Stdin
	file, _ := os.Open(`data.txt`)
	defer file.Close()
	data := readData(file)

	fmt.Printf("%s\n", data)
}

func readData(r io.Reader) []string {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		panic(err)
	}
	return strings.Fields(string(b))
}
