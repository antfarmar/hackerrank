package main

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

var (
	//fileIn  = os.Stdin
	fileOut   = os.Stdout
	fileIn, _ = os.Open(`data\sherlockarray.txt`)
	rw        = bufio.NewReadWriter(bufio.NewReader(fileIn), bufio.NewWriter(fileOut))
)

// MAIN ==================================================

func main() {
	defer fileOut.Close()
	defer fileIn.Close()
	defer rw.Flush()

	T := atois(strings.Fields(readline()))[0]

	for t := 0; t < T; t++ {
		_ = readline()
		A := atois(strings.Fields(readline()))
		out([]string{"NO\n", "YES\n"}[hasSum(A)])
	}
}

// SOLUTION ==============================================

func hasSum(A []int) (has int) {
	L := make([]int, len(A))
	R := make([]int, len(A))

	L[0] = A[0]
	for i := 1; i < len(A); i++ {
		L[i] = A[i] + L[i-1]
	}

	R[len(A)-1] = A[len(A)-1]
	for i := len(A) - 2; i >= 0; i-- {
		R[i] = A[i] + R[i+1]
	}

	//fmt.Printf("A:%v\nL:%v\nR:%v\n", A, L, R)

	for i := 0; i < len(L); i++ {
		if L[i] == R[i] {
			has = 1
			break
		}
	}

	return
}

// IO ====================================================

func out(s string) {
	_, err := rw.WriteString(s)
	check(err)
}

func readline() (line string) {
	line, err := rw.ReadString('\n')
	line = strings.TrimSpace(line)
	check(err)
	return
}

func atois(S []string) []int {
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
