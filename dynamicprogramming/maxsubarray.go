package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var (
	//fileIn  = os.Stdin
	fileOut    = os.Stdout
	fileIn, _  = os.Open(`data\maxsubarray.txt`)
	rw         = bufio.NewReadWriter(bufio.NewReader(fileIn), bufio.NewWriter(fileOut))
	bytesWrote = 0
)

// MAIN ==================================================

func main() {
	defer report()
	defer fileOut.Close()
	defer fileIn.Close()
	defer rw.Flush()

	T := atois(readfields())[0]

	for t := 0; t < T; t++ {
		_ = readline() //skip N
		array := atois(readfields())
		maxSumArr, maxSumSeq := solve(array)
		out(fmt.Sprintln(maxSumArr, maxSumSeq))
	}
}

// SOLUTION ==============================================

func solve(A []int) (maxSumArr, maxSumSeq int) {
	maxSumArr = maxSubarrSum(A)
	maxSumSeq = maxSubseqSum(A)
	if maxSumArr == 0 {
		maxSumArr = maxSumSeq
	}
	return
}

func maxSubarrSum(A []int) (maxSum int) {
	currSum := 0

	for _, n := range A {
		sum := currSum + n
		if sum > 0 {
			currSum = sum
		} else {
			currSum = 0
		}

		if currSum > maxSum {
			maxSum = currSum
		}
	}
	return
}

func maxSubseqSum(A []int) (maxSum int) {
	maxInList := -int(^uint(0) >> 1)
	for _, n := range A {
		if n > 0 {
			maxSum += n
		}
		if n > maxInList {
			maxInList = n
		}
	}
	if maxSum == 0 {
		maxSum = maxInList
	}
	return
}

// IO ====================================================

func out(s string) {
	n, err := rw.WriteString(s)
	check(err)
	bytesWritten += n
}

func report() {
	println("<Wrote", bytesWritten, "bytes to fileOut>")
}

func readline() (line string) {
	line, err := rw.ReadString('\n')
	line = strings.TrimSpace(line)
	check(err)
	return
}

func readfields() (line []string) {
	return strings.Fields(readline())
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
