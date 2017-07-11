package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

var re = regexp.MustCompile("[0-9]+")

func main() {
	file, _ := os.Open(`..\data.txt`)
	//file := os.Stdin
	in := bufio.NewScanner(file)
	in.Scan()
	for in.Scan() {
		line := in.Text()
		m := re.FindAllString(line, -1)
		fmt.Printf("CountryCode=%v,LocalAreaCode=%v,Number=%v\n", m[0], m[1], m[2])
	}
}
