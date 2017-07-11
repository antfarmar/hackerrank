package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Card struct {
	value int
	under *Card
}

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

func LIS(A []int) (result []int) {
	TOPS := make([]*Card, 0, len(A))

	// Patience sort: sort A into piles == LIS of A
	for _, n := range A {

		// Binary search to find and return the smallest index i in [0, n) at which f(i) is true
		j := sort.Search(len(TOPS), func(i int) bool { return TOPS[i].value >= n })

		card := &Card{n, nil}
		if j != 0 {
			card.under = TOPS[j-1]
		}

		if j != len(TOPS) {
			TOPS[j] = card
		} else {
			TOPS = append(TOPS, card)
		}
	}

	if len(TOPS) == 0 {
		return
	}

	// Build the sequence (reversed) by following pointers
	for pile := TOPS[len(TOPS)-1]; pile != nil; pile = pile.under {
		result = append(result, pile.value)
	}

	// Reverse the result
	for i := 0; i < len(result)/2; i++ {
		result[i], result[len(result)-i-1] = result[len(result)-i-1], result[i]
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
