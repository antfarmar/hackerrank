package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var F map[int64]bool
var Fmax int64

func isFibo(n int64) bool {
	_, ok := F[n]
	return ok
}

func init() {
	F = make(map[int64]bool, 100000)
	a := int64(0)
	b := int64(1)
	for a <= 10000000000 {
		a, b = b, a+b
		F[a] = true
	}
}

func main() {
	input := NewInputParser(os.Stdin)
	var T int
	input.Scan(&T)

	for t := 0; t < T; t++ {
		n := input.ParseInt()
		if isFibo(n) {
			fmt.Println("IsFibo")
		} else {
			fmt.Println("IsNotFibo")
		}
	}
}

/* Prints an array without the brakets */
func printArray(a []int) {
	fmt.Println(strings.Trim(fmt.Sprint(a), "[]"))
}

type InputParser struct {
	*bufio.Reader
}

func NewInputParser(rd io.Reader) *InputParser {
	return &InputParser{bufio.NewReader(rd)}
}

/* Returns a single line of input (without the ending newline) */
func (in *InputParser) ParseString() string {
	s, err := in.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	s = strings.TrimSpace(s)
	return s
}

/* Parses a single integer */
func (in *InputParser) ParseInt() int64 {
	s := in.ParseString()
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(err)
	}
	return i
}

/* Scan for variables (pass in pointers)
   var a int
   var b float
   var c string
   in.Scan(&a, &b, &c)
*/
func (in *InputParser) Scan(a ...interface{}) int {
	s := in.ParseString()
	i, err := fmt.Sscan(s, a...)
	if err != nil {
		panic(err)
	}
	return i
}

/* Parses an array of integers separated by space */
func (in *InputParser) ParseIntArray(capacity int) []int {
	a := make([]int, 0, capacity)
	s := in.ParseString()
	for _, w := range strings.Split(s, " ") {
		n, err := strconv.Atoi(w)
		if err != nil {
			panic(err)
		}
		a = append(a, n)
	}
	return a
}
