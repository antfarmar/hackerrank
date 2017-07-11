// findhackerrank
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

func main() {
	b, _ := ioutil.ReadAll(os.Stdin)
	bb := strings.Split(strings.TrimSpace(string(b)), "\n")
	re := regexp.MustCompile("hackerrank")

	for _, line := range bb[1:] {
		line := strings.TrimSpace(line)
		ind := re.FindAllStringIndex(line, -1)
		ans := interpret(ind, len(line))
		fmt.Fprintf(os.Stderr, "%2v %s %v\n", ans, line, ind)
		fmt.Println(ans)
	}
}

func interpret(indices [][]int, n int) (ans int) {
	m := len(indices) - 1

	if m < 0 {
		return -1
	}
	if indices[0][0] == 0 { //starts 1
		ans += 1
	}
	if indices[m][1] == n { //ends 2
		ans += 2
	}
	if ans == 0 { // no match -1
		ans--
	}
	return ans % 3
}
