package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var (
		n    int
		m    [][]int
		s    string
		pass bool
	)

	reader := bufio.NewReader(os.Stdin)
	s, _ = reader.ReadString('\n')
	n, _ = strconv.Atoi(strings.TrimSpace(s))
	m = make([][]int, n)

	for i := 0; i < n; i++ {
		s, _ = reader.ReadString('\n')
		s = strings.TrimSpace(s)
		m[i] = make([]int, n)
		for j, runeValue := range s {
			// fmt.Println(i, j, string(runeValue))
			m[i][j], _ = strconv.Atoi(string(runeValue))
		}
	}

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
				fmt.Print(m[i][j])
			}
		}
		fmt.Println()
	}
}
