package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	file, err := os.Open(`.\data\palindromeindex.txt`)
	check(err)
	defer file.Close()

	process(solution, file)
	//process(solution, os.Stdin)
}

func solution(input []string) (output []string) {
	for _, s := range input[1:] {
		ans := fmt.Sprintln(palIndex(s))
		output = append(output, ans)
	}
	return output
}

func palIndex(s string) int {
	index, n := -1, len(s)
	for i := 0; i < n/2; i++ {
		end := n - i - 1
		if s[i] != s[end] {
			if s[i] == s[end-1] && s[i+1] == s[end-2] {
				return end
			} else {
				return i
			}
		}
	}
	return index
}

// IO UTILITIES =======================================
func process(f func([]string) []string, r io.Reader) {
	b, err := ioutil.ReadAll(r)
	check(err)
	input := strings.Fields(string(b))
	output := fmt.Sprint(f(input))
	fmt.Println(output[1 : len(output)-1])
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
