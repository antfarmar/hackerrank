package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	byteslice, _ := ioutil.ReadFile(`..\data.txt`)
	//byteslice, _ := ioutil.ReadAll(os.Stdin)
	mat := bytes.Fields(byteslice)[1:]
	n := len(mat)

	if n == 1 {
		fmt.Println(string(mat[0][0]))
		os.Exit(0)
	}

	ans := clone(mat)
	for r, row := range mat[1 : n-1] { // row = []byte
		for c, col := range row[1 : n-1] { // col = byte
			lessLR := mat[r+1][c] < col && col > mat[r+1][c+2]
			lessUD := mat[r][c+1] < col && col > mat[r+2][c+1]
			if lessLR && lessUD {
				ans[r+1][c+1] = 'X'
			}
		}
	}
	fmt.Print(sstring(ans))
}

func clone(b [][]byte) (c [][]byte) {
	c = make([][]byte, len(b))
	for i := range b {
		c[i] = make([]byte, len(b[i]))
		copy(c[i], b[i])
	}
	return
}

func sstring(sbs [][]byte) (out string) {
	for _, bs := range sbs {
		out += fmt.Sprintf("%s\n", bs)
	}
	return
}
