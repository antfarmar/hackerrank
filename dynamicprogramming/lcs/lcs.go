package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	//fileIn  = os.Stdin
	fileOut      = os.Stdout
	fileIn, _    = os.Open(`data\lcs.txt`)
	rw           = bufio.NewReadWriter(bufio.NewReader(fileIn), bufio.NewWriter(fileOut))
	bytesWritten = 0
)

// MAIN ==================================================

func main() {
	defer report()
	defer fileOut.Close()
	defer fileIn.Close()
	defer rw.Flush()

	_ = readline() // skip N,M
	A := atois(readfields())
	B := atois(readfields())

	write(strings.Trim(fmt.Sprint(LCS(A, B)), "[]") + "\n")
}

// SOLUTION ==============================================

// DP LCS: O(|A|x|B|)
func LCS(A, B []int) (SEQ []int) {
	var LA, LB = len(A), len(B)
	var LENGTHS [][]int = makeMemo(LA, LB)

	// For each substring of A...
	for a := 1; a <= LA; a++ {

		// For each substring of B...
		for b := 1; b <= LB; b++ {

			switch { // Keep a running +1 tab of matches
			case A[a-1] == B[b-1]:
				LENGTHS[a][b] = LENGTHS[a-1][b-1] + 1
			case LENGTHS[a][b-1] >= LENGTHS[a-1][b]:
				LENGTHS[a][b] = LENGTHS[a][b-1] + 0
			default:
				LENGTHS[a][b] = LENGTHS[a-1][b] + 0

			}
		}
	}

	SEQ = decodeLCS(A, LENGTHS)
	return
}

func decodeLCS(A []int, MEM [][]int) (SEQ []int) {
	I, J := len(MEM)-1, len(MEM[0])-1
	SEQ = make([]int, 0, MEM[I][J])

	// Walk backwards through the memo matrix looking for cells
	// with length values > those found above and to the left
	for i, j := I, J; i != 0 && j != 0; {
		switch {
		case MEM[i][j] == MEM[i-1][j]:
			i--
		case MEM[i][j] == MEM[i][j-1]:
			j--
		default: // a "match" index, so add it
			SEQ = append(SEQ, A[i-1])
			i--
			j--
		}
	}

	// Reverse SEQ 'stack' and return it
	for i, j := 0, len(SEQ)-1; i < j; i, j = i+1, j-1 {
		SEQ[i], SEQ[j] = SEQ[j], SEQ[i]
	}
	return
}

func makeMemo(m, n int) (MN [][]int) {
	MN = make([][]int, m+1)
	for i := range MN {
		MN[i] = make([]int, n+1)
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
