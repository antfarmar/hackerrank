package main

import (
	"bufio"
	"os"
)

func main() {
	process(solution)
}

func solution(input []string) (output []string) {
	output = append(output, "answer")
	return output
}

func process(f func([]string) []string) {
	var (
		input   []string
		scanner = bufio.NewScanner(os.Stdin)
		writer  = bufio.NewWriter(os.Stdout)
	)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	for _, s := range f(input) {
		writer.WriteString(s + "\n")
	}
	writer.Flush()
}
