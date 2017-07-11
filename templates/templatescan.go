package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanLines)
	in.Scan() // skip n
	for in.Scan() {
		s := in.Text()
		fmt.Println(s)
	}
}

/* Or fmt.Scan
func scan(a ...interface{}) {
	if _, err := fmt.Scan(a...); err != nil {
		panic(err)
	}
}
*/
