package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

var (
	re  = regexp.MustCompile("(?i:hackerrank)")
	cnt = 0
)

func main() {
	in := bufio.NewScanner(os.Stdin)
	in.Scan() // don't care # cases
	for in.Scan() {
		s := in.Text()
		if re.MatchString(s) {
			cnt++
		}
	}
	fmt.Println(cnt)
}
