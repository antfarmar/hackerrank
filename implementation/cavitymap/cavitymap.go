package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

func main() {
	byteslice, _ := ioutil.ReadFile(`..\data.txt`)
	//byteslice, _ := ioutil.ReadAll(os.Stdin)
	m := bytes.Fields(byteslice)[1:]
	n := len(m)
	pass := false

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			pass = i > 0 && i < n-1 && j > 0 && j < n-1
			pass = pass && m[i][j] > m[i-1][j]
			pass = pass && m[i][j] > m[i][j-1]
			pass = pass && m[i][j] > m[i][j+1]
			pass = pass && m[i][j] > m[i+1][j]
			if pass {
				fmt.Print("X")
			} else {
				fmt.Print(string(m[i][j]))
			}
		}
		fmt.Print("\n")
	}
}
