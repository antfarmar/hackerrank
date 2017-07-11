package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

var re = regexp.MustCompile(`^([a-z]{0,3}[0-9]{2,8}[A-Z]{3,}$)`)

func main() {
	//file := os.Stdin
	file, _ := os.Open(`..\data.txt`)
	in := bufio.NewScanner(file)
	in.Scan()
	for in.Scan() {
		if parse(in.Text()) {
			fmt.Println("VALID")
		} else {
			fmt.Println("INVALID")
		}
	}
}

func parse(str string) (mat bool) {
	fnd := re.FindString(str)
	mat = re.MatchString(str)
	fmt.Fprintf(os.Stderr, "%35v\t%v\t%v\n", str, mat, fnd)
	return
}
