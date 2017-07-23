package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var (
	//fileIn        = os.Stdin
	testData        = `C:\Users\maf\go\src\hackerrank\strings\caesarcipher\input\input00.txt`
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

	readline() //skip N
	S := readline()
	K := atois(readfields())[0]
	write(fmt.Sprintf("%v\n", solution(S, K)))
}

// SOLUTION ==============================================

func solution(s string, rot int) (cipher string) {
	for _, arune := range s {
		if unicode.IsLetter(arune) {
			arune = rotate(arune, rune(rot))
		}
		cipher += string(arune)
	}
	return
}

func rotate(r, amt rune) (new rune) {
	start := 'A'
	if unicode.IsLower(r) {
		start = 'a'
	}
	new = ((r - start + amt) % 26) + start
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
