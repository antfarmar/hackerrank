package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	data := scanIntoSlice(01, 01)
	ints := []int{}
	cnt, tot := 0, 0

	for _, d := range data {
		i, _ := strconv.Atoi(d)
		ints = append(ints, i)
	}

	sort.Sort(sort.IntSlice(ints[2:]))

	for _, cst := range ints[2:] {
		tot += cst
		if tot < ints[1] {
			cnt++
		} else {
			break
		}
	}

	fmt.Println(cnt)
	debug(ints)
}

// IO =========================================================================
func scanIntoSlice(stdin, linesWords int) []string {
	var file *os.File
	if stdin == 0 {
		file = os.Stdin
	} else {
		var err error
		file, err = os.Open(`data\markandtoys.txt`)
		check(err)
		defer file.Close()
	}
	scanner := bufio.NewScanner(file)
	if linesWords == 0 {
		scanner.Split(bufio.ScanLines)
	} else {
		scanner.Split(bufio.ScanWords)
	}
	input := []string{}
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	check(scanner.Err())
	return input
}

func sprint(ss []string, name string) {
	if name == "" {
		name = "SLICE"
	}
	fmt.Println(sstring(ss, name))
}

func sstring(ss []string, name string) (out string) {
	out += fmt.Sprintf("%v: %q\n", name, ss)
	for i, s := range ss {
		out += fmt.Sprintf("[%v]: %v\n", i, s)
	}
	return
}

func debug(args ...interface{}) {
	fmt.Fprintln(os.Stderr, args)
}

func check(err error) {
	if err != nil {
		panic(err)
		os.Exit(1)
	}
}
