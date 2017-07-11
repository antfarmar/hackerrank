package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

// JUST NUMBERS `[+-]?(?:0|[1-9][0-9]*)(?:\.[0-9]+)?`
var re = regexp.MustCompile(
	`\([+-]?((90(\.0+)?)|((\d|[1-8]\d)(\.\d+)?)), [+-]?((180(\.0+)?)|((\d?\d|1[0-7]\d)(\.\d+)?))\)`)

func main() {
	//file := os.Stdin
	file, _ := os.Open(`..\data.txt`)
	in := bufio.NewScanner(file)
	in.Scan()
	for in.Scan() {
		if parseLatLong(in.Text()) {
			fmt.Println("Valid")
		} else {
			fmt.Println("Invalid")
		}
	}
}

func parseLatLong(str string) (mat bool) {
	fnd := re.FindString(str)
	mat = re.MatchString(str)
	fmt.Fprintf(os.Stderr, "%35v\t%v\t%v\n", str, mat, fnd)
	return
}
