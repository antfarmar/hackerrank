package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

var (
	//fileIn       = os.Stdin
	fileOut      = os.Stdout
	fileIn, _    = os.Open(`data\twoarrays.txt`)
	rw           = bufio.NewReadWriter(bufio.NewReader(fileIn), bufio.NewWriter(fileOut))
	bytesWritten = 0
)

// MAIN ==================================================

func main() {
	defer report()
	defer fileOut.Close()
	defer fileIn.Close()
	defer rw.Flush()

	for T := atois(readfields())[0]; T > 0; T-- {
		NK := atois(readfields())
		K := NK[1]
		A := atois(readfields())
		B := atois(readfields())
		write([]string{"YES", "NO"}[greaterThanK(K, A, B)] + "\n")
	}
}

// SOLUTION ==============================================

func greaterThanK(K int, A, B []int) (nope int) {
	sort.Ints(A)
	sort.Sort(sort.Reverse(sort.IntSlice(B)))

	for i := 0; i < len(A); i++ {
		if A[i]+B[i] < K {
			nope = 1
			break
		}
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
