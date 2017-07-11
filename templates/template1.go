package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	//data := readData(`stdin`)
	data := readData(`..\data.txt`)
	fmt.Printf("%s\n", data)
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
