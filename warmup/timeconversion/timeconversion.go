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
	testData        = `C:\Users\maf\go\src\hackerrank\warmup\timeconversion\input\input00.txt`
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

	T := readfields()[0]
	write(fmt.Sprintf("%v\n", solution(T)))
}

// SOLUTION ==============================================

func solution(t12 string) (t24 string) {
	AMPM := t12[8:]
	hour, err := strconv.Atoi(t12[:2])
	check(err)
	hour += 12
	if AMPM == "AM" {
		hour %= 12
	} else if hour == 24 { //12:00:00PM
		hour = 12
	}
	t24 = fmt.Sprintf("%02v"+t12[2:8], hour)
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
