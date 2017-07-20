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
	fileIn, _    = os.Open(`data\knapsack.txt`)
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

	for t := 0; t < T; t++ {
		K := atois(readfields())[1]
		A := atois(readfields())
		write(fmt.Sprintln(maximizeSum(A, K)))
	}
}

// SOLUTION ==============================================

func maximizeSum(A []int, target int) (ans int) {
	MEM := make([]int, target+1) // holds closest subsums to subtargets in [1..target]
	SUB := make([]int, len(A))   // holds current subsums computed per subtarget

	// For each subtarget...
	for T := 1; T <= target; T++ {

		// For each subsum instance...
		for i := 0; i < len(A); i++ {
			if T >= A[i] {
				SUB[i] = MEM[T-A[i]] + A[i]
			} else {
				SUB[i] = 0
			}
		}
		// Memo the closest subsum to subtarget...
		for _, subsum := range SUB {
			if subsum > MEM[T] {
				MEM[T] = subsum
			}
		}
	}

	// Return the closest sum <= target found
	ans = MEM[target]
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
