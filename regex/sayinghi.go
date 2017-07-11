package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

var re = regexp.MustCompile(`^(?i:hi [^d])`)

func main() {
	//file := os.Stdin
	file, _ := os.Open(`..\data.txt`)
	in := bufio.NewScanner(file)
	in.Scan()
	for in.Scan() {
		line := in.Text()
		if parse(line) {
			fmt.Println(line)
		}
	}
}

func parse(str string) (mat bool) {
	fnd := re.FindString(str)
	mat = re.MatchString(str)
	fmt.Fprintf(os.Stderr, "%v\t%v\t%v\n", mat, fnd, str)
	return
}
