package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"regexp"
)

var (
	buff bytes.Buffer
	idRE = regexp.MustCompile(`"question-summary-([0-9]+)`)
	quRE = regexp.MustCompile(`"question-hyperlink">(.*)</a>`)
	tmRE = regexp.MustCompile(`"relativetime">(.*)</span>`)
)

func main() {
	file, _ := os.Open(`..\data.txt`)
	//file := os.Stdin
	in := bufio.NewScanner(file)
	for in.Scan() {
		parse(in.Text())
	}
	fmt.Println(buff.String())
}

func parse(line string) {
	if idRE.MatchString(line) {
		buff.WriteString(idRE.FindStringSubmatch(line)[1] + ";")
	}
	if quRE.MatchString(line) {
		buff.WriteString(quRE.FindStringSubmatch(line)[1] + ";")
	}
	if tmRE.MatchString(line) {
		buff.WriteString(tmRE.FindStringSubmatch(line)[1] + "\n")
	}
	return
}
