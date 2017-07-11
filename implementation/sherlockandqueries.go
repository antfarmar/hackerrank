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
	fileIn, _    = os.Open(`.\data\sherlockandqueries.txt`)
	rw           = bufio.NewReadWriter(bufio.NewReader(fileIn), bufio.NewWriter(fileOut))
	bytesWritten = 0
)

// MAIN ==================================================

func main() {
	defer report()
	defer fileOut.Close()
	defer fileIn.Close()
	defer rw.Flush()

	_ = readline() // skip M,N
	A := atois(readfields())
	B := atois(readfields())
	C := atois(readfields())

	write(strings.Trim(fmt.Sprint(solve(A, B, C)), "[]") + "\n")

}

// SOLUTION ==============================================
/*
for i = 1 to M do
    for j = 1 to N do
        if j % B[i] == 0 then
            A[j] = A[j] * C[i]
        endif
    end do
end do
*/
func solve(A, B, C []int) []int {
	M, N := len(B), len(A)
	mod := uint64(10e9 + 7)

	prods := make(map[int]int, M)

	for i := 0; i < M; i++ {
		if prods[B[i]] == 0 {
			prods[B[i]] = C[i]
		} else {
			prd := prods[B[i]]
			prods[B[i]] = int((uint64(C[i]) * uint64(prd)) % mod)
		}
	}

	for div, prd := range prods {
		for j := div - 1; j < N; j += div {
			A[j] = int((uint64(A[j]) * uint64(prd)) % mod)
		}
	}
	return A
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
