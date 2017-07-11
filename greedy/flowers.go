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
	//fileIn       = os.Stdin
	fileOut      = os.Stdout
	fileIn, _    = os.Open(`data\flowers.txt`)
	rw           = bufio.NewReadWriter(bufio.NewReader(fileIn), bufio.NewWriter(fileOut))
	bytesWritten = 0
)

// MAIN ==================================================

func main() {
	defer report()
	defer fileOut.Close()
	defer fileIn.Close()
	defer rw.Flush()

	NK := atois(readfields())
	N, K := NK[0], NK[1]
	COSTS := atois(readfields())
	sort.Sort(sort.Reverse(sort.IntSlice(COSTS)))
	write(fmt.Sprintln(minCost(COSTS, N, K)))
}

// SOLUTION ==============================================
// Each person takes turns buying the highest costing flower
// available first, increasing the factor by 1 each round.

func minCost(COSTS []int, N, K int) (total int) {
	buys, factor := 1, 1
	for _, cost := range COSTS {
		total += cost * factor
		buys++
		if buys > K {
			buys = 1
			factor++
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
