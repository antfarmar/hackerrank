package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	//fileIn        = os.Stdin
	testData        = `C:\Users\maf\go\src\hackerrank\implementation\appleandorange\input\input00.txt`
	fileIn, openErr = os.Open(testData)
	fileOut         = os.Stdout
	rw              = bufio.NewReadWriter(bufio.NewReader(fileIn), bufio.NewWriter(fileOut))
	bytesWritten    = 0
)

// MAIN ==================================================

func main() {
	check(openErr)
	defer report()
	defer fileOut.Close()
	defer fileIn.Close()
	defer rw.Flush()

	ST := atois(readfields())
	AB := atois(readfields())
	readfields() //skip MN
	M := atois(readfields())
	N := atois(readfields())
	s, t := ST[0], ST[1]
	a, b := AB[0], AB[1]
	write(fmt.Sprintf("%v\n", solution(a, s, t, M))) //apples
	write(fmt.Sprintf("%v\n", solution(b, s, t, N))) //oranges
}

// SOLUTION ==============================================

func solution(c, s, t int, I []int) (hits int) {
	for _, i := range I {
		if c+i >= s && c+i <= t {
			hits++
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
	fmt.Fprintf(os.Stderr, "<wrote %v bytes>", bytesWritten)
}

func readline() (line string) {
	line, err := rw.ReadString('\n')
	check(err)
	line = strings.TrimSpace(line)
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
			fmt.Fprintln(os.Stderr, "<EOF reached>")
		} else {
			log.Fatal(err)
		}
	}
}
