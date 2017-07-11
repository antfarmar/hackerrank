package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	file, err := os.Open(`data.txt`)
	check(err)
	defer file.Close()

	process(solution, file)
	//process(solution, os.Stdin)
}

func solution(input []string) (output []string) {
	output = input
	debug("Debug:", fmt.Sprintf("%b", 251976))
	return output
}

// IO UTILITIES =======================================
func process(f func([]string) []string, r io.Reader) {
	b, err := ioutil.ReadAll(r)
	check(err)
	input := strings.Fields(string(b))
	for _, s := range f(input) {
		fmt.Println(s)
	}
}

func check(err error) {
	if err != nil {
		panic(err)
		os.Exit(1)
	}
}

func debug(args ...interface{}) {
	fmt.Fprintln(os.Stderr, args)
}
