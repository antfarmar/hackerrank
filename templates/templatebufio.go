package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func solution(input []string) (output []string) {
	for _, sample := range input {
		output = append(output, f(sample))
	}
	return output
}

func f(s string) string {
	debug("Debug:", fmt.Sprintf("%b", 251976))
	return s
}

// MAIN ==================================================

func main() {
	file, err := os.Open(`data.txt`)
	check(err)
	defer file.Close()

	process(solution, file)
	//process(solution, os.Stdin)
}

// IO UTILITIES ==========================================

func process(f func([]string) []string, file io.Reader) {
	var (
		input   []string
		scanner = bufio.NewScanner(file) // default is ScanLine
		writer  = bufio.NewWriter(os.Stdout)
	)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		input = append(input, scanner.Text()) // scans lines
	}
	check(scanner.Err())

	for _, s := range f(input) {
		writer.WriteString(s + "\n")
	}
	writer.Flush()
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
