package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	//fileIn       = os.Stdin
	fileOut      = os.Stdout
	fileIn, _    = os.Open(`.\data\makeitanagram.txt`)
	rw           = bufio.NewReadWriter(bufio.NewReader(fileIn), bufio.NewWriter(fileOut))
	bytesWritten = 0
)

// MAIN ==================================================

func main() {
	defer report()
	defer fileOut.Close()
	defer fileIn.Close()
	defer rw.Flush()

	A, B := readline(), readline()
	write(fmt.Sprintf("%d\n", makeItAnagram(A, B)))
}

// SOLUTION ==============================================

func makeItAnagram(A, B string) (minDelete int) {
	chrCnt := make(map[rune]int, 26)

	for _, c := range A {
		chrCnt[c]++
	}

	for _, c := range B {
		chrCnt[c]--
	}

	// Count differences between A and B
	for _, val := range chrCnt {
		minDelete += abs(val)
	}

	return
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// IO ====================================================

func write(s string) {
	n, err := rw.WriteString(s)
	check(err)
	bytesWritten += n
}

func report() {
	println("<wrote", bytesWritten, "bytes>")
}

func readline() (line string) {
	line, err := rw.ReadString('\n')
	line = strings.TrimSpace(line)
	check(err)
	return
}

func readfields() (fields []string) {
	line, err := rw.ReadString('\n')
	fields = strings.Fields(line)
	check(err)
	return
}

func atois(S []string) (I []int) {
	I = make([]int, len(S))
	for i, a := range S {
		n, err := strconv.Atoi(a)
		check(err)
		I[i] = n
	}
	return
}

func check(err error) {
	if err != nil {
		if err.Error() == "EOF" {
			println("<EOF reached>")
		} else {
			panic(err)
		}
	}
}
