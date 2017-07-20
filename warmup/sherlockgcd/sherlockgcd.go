package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	//file := os.Stdin
	file, _ := os.Open(`.\data\sherlockgcd.txt`)
	rawBytes, _ := ioutil.ReadAll(file)
	cases := strings.Fields(string(rawBytes))[1:]

	// For each test case..
	for i := 0; i < len(cases); {
		length, _ := strconv.Atoi(cases[i])
		div, ans := 0, "NO"

		// Parse numbers
		for j := 0; j < length; j++ {
			n, _ := strconv.Atoi(cases[i+j+1])

			if div == 0 {
				div = n
			} else {
				div = gcd(div, n)
			}

			if div == 1 {
				ans = "YES"
				break
			}
		}

		// Print the answer for this case
		fmt.Println(ans)

		// Continue to next test case set
		i += length + 1
	}
}

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}
