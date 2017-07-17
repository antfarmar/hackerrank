package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

var (
	//fileIn        = os.Stdin
	testData        = `C:\Users\maf\go\src\hackerrank\warmup\diagonaldifference\input\input00.txt`
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

	N := atois(readfields())[0]
	write(fmt.Sprintf("%v\n", solution(N)))
}

// SOLUTION ==============================================

func solution(n int) (absDiff int) {
	var diagSum1, diagSum2 int
	for i := 0; i < n; i++ {
		ROW := atois(readfields())
		diagSum1 += ROW[i]
		diagSum2 += ROW[n-i-1]
	}
	absDiff = int(math.Abs(float64(diagSum1 - diagSum2)))
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
