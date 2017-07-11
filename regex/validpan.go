package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

//<char><char><char><char><char><digit><digit><digit><digit><char>
var re = regexp.MustCompile("[A-Z]{5}[0-9]{4}[A-Z]")

func main() {
	in := bufio.NewScanner(os.Stdin)
	in.Scan() // don't care # cases
	for in.Scan() {
		s := in.Text()
		switch re.MatchString(s) {
		case true:
			fmt.Println("YES")
		case false:
			fmt.Println("NO")
		}
	}
}
