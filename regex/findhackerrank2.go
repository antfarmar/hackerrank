package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const magic = "hackerrank"

func main() {
	in := bufio.NewScanner(os.Stdin)
	in.Scan() // skip N
	for in.Scan() {
		s := in.Text()
		switch pre, suf := strings.HasPrefix(s, magic), strings.HasSuffix(s, magic); {
		case pre && suf:
			fmt.Println(0)
		case pre:
			fmt.Println(1)
		case suf:
			fmt.Println(2)
		default:
			fmt.Println(-1)
		}
	}
}
