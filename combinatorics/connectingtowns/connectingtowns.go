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
	fileIn, _    = os.Open(`data\connectingtowns.txt`)
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
		_ = readline() // skip N
		totRoutes := atois(readfields())
		write(fmt.Sprintln(totRoutesT1toTn(totRoutes)))
	}
}

// SOLUTION ==============================================

func totRoutesT1toTn(totRoutes []int) (total int) {
	total = 1
	for _, n := range totRoutes {
		total *= n
		total %= 1234567
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
