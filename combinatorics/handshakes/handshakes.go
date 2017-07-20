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
	fileIn, _    = os.Open(`data\handshakes.txt`)
	rw           = bufio.NewReadWriter(bufio.NewReader(fileIn), bufio.NewWriter(fileOut))
	bytesWritten = 0
)

// MAIN ==================================================

func main() {
	defer report()
	defer fileOut.Close()
	defer fileIn.Close()
	defer rw.Flush()

	T := atois(readfields())[0]

	for i := 0; i < T; i++ {
		totPeople := atois(readfields())[0]
		write(fmt.Sprintln(handshakes(totPeople)))
	}
}

// SOLUTION ==============================================

func handshakes(totPeople int) (shakes int) {
	shakes = (totPeople * (totPeople - 1)) / 2 // n choose 2
	return
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
