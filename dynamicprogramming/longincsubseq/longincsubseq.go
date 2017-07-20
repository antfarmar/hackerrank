package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var (
	//fileIn  = os.Stdin
	fileOut      = os.Stdout
	fileIn, _    = os.Open(`data\longincsubseq.txt`)
	rw           = bufio.NewReadWriter(bufio.NewReader(fileIn), bufio.NewWriter(fileOut))
	bytesWritten = 0
)

// MAIN ==================================================

func main() {
	defer report()
	defer fileOut.Close()
	defer fileIn.Close()
	defer rw.Flush()

	N := atois(readfields())[0]
	A := make([]int, N)
	B := make([]int, N)

	for n := 0; n < N; n++ {
		A[n] = atois(readfields())[0]
		B[n] = A[n]
	}

	sort.Ints(B)
	fmt.Println(A, B)
	write(fmt.Sprintln(LCS(A, B)))
}

// SOLUTION ==============================================

// LCS with repetition (not strictly increasing)
// O(n^2) (Patience sort is O(n lg n))
func LCS(A, B []int) (length int) {
	var M, N = len(A), len(B)
	var MEM [][]int = makeMemo(M, N)

	for i := 1; i <= M; i++ {
		for j := 1; j <= N; j++ {
			switch {
			case A[i-1] == B[j-1]:
				MEM[i][j] = MEM[i-1][j-1] + 1
			case MEM[i][j-1] >= MEM[i-1][j]:
				MEM[i][j] = MEM[i][j-1] + 0
			default:
				MEM[i][j] = MEM[i-1][j] + 0
			}
		}
	}

	printlcs(A, MEM, M, N)
	fmt.Println()

	length = MEM[M][N]
	return
}

func makeMemo(m, n int) (memo [][]int) {
	memo = make([][]int, m+1)
	for i := range memo {
		memo[i] = make([]int, n+1)
	}
	return
}

func printlcs(s []int, r [][]int, i, j int) {
	if i == 0 || j == 0 {
		return
	}
	switch {
	case r[i][j] > r[i-1][j] && r[i][j] > r[i][j-1]:
		printlcs(s, r, i-1, j-1)
		fmt.Printf("%v ", s[i-1])
	case r[i][j-1] >= r[i-1][j]:
		printlcs(s, r, i, j-1)
	default:
		printlcs(s, r, i-1, j)
	}
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
			println("<EOF reached>")
		} else {
			panic(err)
		}
	}
}
