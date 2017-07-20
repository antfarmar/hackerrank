// insertionsort1
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	b, err := ioutil.ReadAll(os.Stdin)
	//b, err := ioutil.ReadFile(`..\data.txt`)
	if err != nil {
		panic(err)
	}
	data := strings.Fields(string(b))
	nums := parseInts(data[1:])
	insertionSort(nums)
}

func insertionSort(array []int) {
	for i := 1; i < len(array); i++ {
		for j := i; j > 0 && array[j] < array[j-1]; j-- {
			//array[j], array[j-1] = array[j-1], array[j]
			temp := array[j]
			array[j] = array[j-1]
			printArray(array)
			array[j-1] = temp
		}
	}
	printArray(array)
}

func printArray(a []int) {
	s := fmt.Sprint(a)
	s = strings.Trim(s, "[]")
	fmt.Println(s)
}

func parseInts(ss []string) (is []int) {
	is = make([]int, len(ss))
	for n, s := range ss {
		is[n], _ = strconv.Atoi(s)
	}
	return
}
