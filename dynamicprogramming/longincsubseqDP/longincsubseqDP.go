package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
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

	for n := 0; n < N; n++ {
		A[n] = atois(readfields())[0]
	}

	fmt.Println(A)
	write(fmt.Sprintln(LIS(A)))
}

// SOLUTION ==============================================

// DP LIS (Strictly increasing) O(n log n)
func LIS(A []int) (SEQ []int) {
	N := len(A)
	IDX := make([]int, N)
	MEM := make([]int, N+1)
	length := 0

	for i := 0; i < N; i++ {
		lo := 1
		hi := length
		for lo <= hi {
			mid := int(math.Ceil(float64(lo+hi) / 2))
			if A[MEM[mid]] < A[i] {
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		}

		newLength := lo
		IDX[i] = MEM[newLength-1]
		MEM[newLength] = i

		if newLength > length {
			length = newLength
		}
	}

	SEQ = make([]int, length)
	k := MEM[length]
	for i := length - 1; i >= 0; i-- {
		SEQ[i] = A[k]
		k = IDX[k]
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
