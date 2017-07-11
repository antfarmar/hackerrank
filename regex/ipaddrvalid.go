package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var (
	ipv4_re = join(`(\d{1,3})`, `\.`, 4)
	ipv6_re = join(`[[:xdigit:]]{1,4}`, `:`, 8) // xdigit ≡ [0-9A-Fa-f]
	ipv4RE  = regexp.MustCompile(`^` + ipv4_re + `$`)
	ipv6RE  = regexp.MustCompile(`^` + ipv6_re + `$`)
)

func main() {
	lines := scanIntoSlice(01, 0)[1:]
	for _, ln := range lines {
		fmt.Println(parseIP(ln))
	}
}

func parseIP(ln string) (ans string) {
	ans = "Neither"
	if ipv4RE.MatchString(ln) {
		asm := ipv4RE.FindAllStringSubmatch(ln, -1)
		if verify0to255(asm[0][1:]) {
			ans = "IPv4"
		}
	} else if ipv6RE.MatchString(ln) {
		ans = "IPv6"
	}
	return
}

func verify0to255(sm []string) bool {
	for _, m := range sm {
		n, _ := strconv.Atoi(m)
		if n > 255 {
			return false
		}
	}
	return true
}

// IO =========================================================================
func scanIntoSlice(stdin, linesWords int) []string {
	var file *os.File
	if stdin == 0 {
		file = os.Stdin
	} else {
		var err error
		file, err = os.Open(`..\data.txt`)
		check(err)
		defer file.Close()
	}
	scanner := bufio.NewScanner(file)
	if linesWords == 0 {
		scanner.Split(bufio.ScanLines)
	} else {
		scanner.Split(bufio.ScanWords)
	}
	input := []string{} //make([]string, 0)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	check(scanner.Err())
	return input
}

func join(a string, sep string, amt int) string {
	n := len(sep)*(amt-1) + len(a)*amt
	b := make([]byte, n)
	bp := copy(b, a)
	for i := 1; i < amt; i++ {
		bp += copy(b[bp:], sep)
		bp += copy(b[bp:], a)
	}
	return string(b)
}

func sstring(ss []string) (out string) {
	out += fmt.Sprintf("SLICE: %q\n", ss)
	for i, s := range ss {
		out += fmt.Sprintf("[%v]: %v\n", i, s)
	}
	return
}

func check(err error) {
	if err != nil {
		panic(err)
		os.Exit(1)
	}
}

func debug(s string) {
	fmt.Fprintln(os.Stderr, s)
}
