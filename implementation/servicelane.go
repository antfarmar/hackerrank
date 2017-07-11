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
	fileIn, _    = os.Open(`.\data\servicelane.txt`)
	rw           = bufio.NewReadWriter(bufio.NewReader(fileIn), bufio.NewWriter(fileOut))
	bytesWritten = 0
)

// MAIN ==================================================

func main() {
	defer report()
	defer fileOut.Close()
	defer fileIn.Close()
	defer rw.Flush()

	NT := atois(readfields())
	T := NT[1]
	WIDTHS := atois(readfields())

	for ; T > 0; T-- {
		IJ := atois(readfields())
		write(fmt.Sprintf("%d\n", largestVehicle(WIDTHS, IJ[0], IJ[1])))
	}

}

// SOLUTION ==============================================

func largestVehicle(widths []int, i, j int) (largest int) {
	largest = 3
	for idx := i; idx <= j; idx++ {
		if widths[idx] < largest {
			largest = widths[idx]
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
