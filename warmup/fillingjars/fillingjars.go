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
	file, _ := os.Open(`.\data\fillingjars.txt`)
	rawBytes, _ := ioutil.ReadAll(file)

	data := strings.Fields(string(rawBytes))
	numJars, _ := strconv.ParseUint(data[0], 10, 64)
	numCandies := uint64(0)
	abk := data[2:]

	// For each case triplet, sum # candies placed in jars
	for i := 0; i <= len(abk)-3; i += 3 {
		a, _ := strconv.Atoi(abk[i+0])
		b, _ := strconv.Atoi(abk[i+1])
		k, _ := strconv.ParseUint(abk[i+2], 10, 64)
		numCandies += k * uint64(1+b-a)
	}

	fmt.Fprintf(os.Stderr, "%v %v %v\n", numJars, numCandies, abk)
	fmt.Printf("%v\n", numCandies/numJars)
}

/************************************
package main

import "fmt"

func main() {
	var n, m int
	var a, b, k int
	fmt.Scanf("%v %v", &n, &m)
	total := 0
	for in := 0; in < m; in++ {
		fmt.Scanf("%v %v %v", &a, &b, &k)
		total += (b - a + 1) * k
	}

	fmt.Println(total / n)
}
***************************************/
