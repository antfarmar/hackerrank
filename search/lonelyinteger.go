package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// SOLUTION ==============================================

func solution(input []string) (output []string) {
	for _, line := range input[1:] {
		sample := strings.Fields(line)
		output = append(output, cntInts(sample))
	}
	return output
}

func cntInts(ss []string) string {
	hash := make(map[string]int, len(ss))
	for _, k := range ss {
		hash[k]++
		//debug(hash)
	}
	for k, _ := range hash {
		if hash[k] == 1 {
			return k
		}
	}
	return "no singles"
}

// MAIN ==================================================

func main() {
	file, err := os.Open(`.\data\lonelyinteger.txt`)
	check(err)
	defer file.Close()

	process(solution, file)
	//process(solution, os.Stdin)
}

// IO UTILITIES ==========================================

func process(soln func([]string) []string, file io.Reader) {
	var (
		input   []string
		scanner = bufio.NewScanner(file)
		writer  = bufio.NewWriter(os.Stdout)
	)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	check(scanner.Err())

	for _, s := range soln(input) {
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
