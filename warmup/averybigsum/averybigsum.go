package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	//fileIn        = os.Stdin
	testData        = `C:\Users\maf\go\src\hackerrank\warmup\averybigsum\a-very-big-sum-testcases\input\input00.txt`
	fileIn, openErr = os.Open(testData)
	fileOut         = os.Stdout
	rw              = bufio.NewReadWriter(bufio.NewReader(fileIn), bufio.NewWriter(fileOut))
	bytesWritten    = 0
)

// MAIN ==================================================

func main() {
	check(openErr)
	defer report()
	defer fileOut.Close()
	defer fileIn.Close()
	defer rw.Flush()

	readfields() //skip N
	A := atois(readfields())
	write(fmt.Sprintf("%v\n", solution(A)))
}

// SOLUTION ==============================================

func solution(A []int) (sum int64) {
	for _, a := range A {
		sum += int64(a)
	}
	return
}

// IO ====================================================

func write(s string) {
	n, err := rw.WriteString(s)
	check(err)
	bytesWritten += n
}

func report() {
	fmt.Fprintf(os.Stderr, "<wrote %v bytes>", bytesWritten)
}

func readline() (line string) {
	line, err := rw.ReadString('\n')
	check(err)
	line = strings.TrimSpace(line)
	return
}

func readfields() (fields []string) {
	line, err := rw.ReadString('\n')
	check(err)
	fields = strings.Fields(line)
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
			fmt.Fprintln(os.Stderr, "<EOF reached>")
		} else {
			log.Fatal(err)
		}
	}
}
