package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var data []string
	data = scanline(data)
	fmt.Println(pangram(data))
}

func pangram(words []string) (ans string) {
	ans = "pangram"
	hash := make(map[rune]bool)

	str := strings.Join(words, "")
	str = strings.ToLower(str)

	for _, r := range str {
		hash[r] = true
	}

	if len(hash) == 26 {
		return ans
	}
	return "not " + ans
}

func scanline(data []string) []string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords) // default is ScanLine
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	return data
}

/*
Sample Input #1
We promptly judged antique ivory buckles for the next prize

Sample Output #1
pangram

Sample Input #2
We promptly judged antique ivory buckles for the prize

Sample Output #2
not pangram
*/
