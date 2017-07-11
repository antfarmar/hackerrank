package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

// MAIN ==================================================

func main() {
	//input,_ := ioutil.ReadAll(os.Stdin)
	input, _ := ioutil.ReadFile(`data\commonchild.txt`)
	words := bytes.Fields(input)
	fmt.Println(commonString(words[0], words[1]))
}

// SOLUTION ==============================================

func commonString(s1, s2 []uint8) (length int) {
	memo := LCS(s1, s2)
	length = memo[len(s1)][len(s2)]
	return
}

func LCS(s1, s2 []byte) [][]int {

	m, n := len(s1), len(s2)
	memo := make([][]int, m+1)

	for i, _ := range memo {
		memo[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			switch {
			case s1[i-1] == s2[j-1]:
				memo[i][j] = memo[i-1][j-1] + 1
			case memo[i][j-1] >= memo[i-1][j]:
				memo[i][j] = memo[i][j-1]
			default:
				memo[i][j] = memo[i-1][j]
			}
		}
	}
	return memo
}
