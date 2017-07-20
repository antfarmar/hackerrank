package main

import (
	"bufio"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

var (
	//fileIn  = os.Stdin
	fileOut   = os.Stdout
	fileIn, _ = os.Open(`data\cutthesticks.txt`)
	rw        = bufio.NewReadWriter(bufio.NewReader(fileIn), bufio.NewWriter(fileOut))
)

// MAIN ==================================================

func main() {
	defer fileOut.Close()
	defer fileIn.Close()
	defer rw.Flush()

	_ = readline() // skip N
	stickLengths := atoi(strings.Fields(readline()))

	ans := cutSticks(stickLengths)
	out(strings.Join(ans, "\n"))
}

// SOLUTION ==============================================

func cutSticks(stickLengths []int) (cutCnts []string) {
	sort.Ints(stickLengths)

	for len(stickLengths) != 0 {
		c := strconv.Itoa(len(stickLengths))
		cutCnts = append(cutCnts, c)
		stickLengths = removeMins(stickLengths, stickLengths[0])
	}
	return
}

func removeMins(A []int, min int) []int {
	for i, n := range A {
		if n > min {
			return A[i:]
		}
	}
	return A[len(A):]
}

// IO ====================================================

func out(s string) {
	n, err := rw.WriteString(s)
	check(err)
	println("\n<wrote", n, "bytes to out file>")
}

func readline() (line string) {
	line, err := rw.ReadString('\n')
	line = strings.TrimSpace(line)
	check(err)
	return
}

func atoi(S []string) []int {
	I := make([]int, len(S))
	for i, a := range S {
		n, err := strconv.Atoi(a)
		check(err)
		I[i] = n
	}
	return I
}

func check(err error) {
	if err != nil && err != io.EOF {
		panic(err)
	}
}
