package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var (
	//fileIn    = os.Stdin
	fileIn, _ = os.Open(`.\data\lonelyinteger.txt`)
	fileOut   = os.Stdout
	rw        = bufio.NewReadWriter(bufio.NewReader(fileIn), bufio.NewWriter(fileOut))
)

// MAIN ==================================================
func main() {
	defer rw.Flush()
	defer fileIn.Close()

	var data string
	var nums []int

	data, _ = rw.ReadString('\n')
	N, _ := strconv.Atoi(strings.TrimSpace(data))
	nums = make([]int, N, N)

	for i := 0; i < N; i++ {
		data, _ := rw.ReadString(' ')
		n, _ := strconv.Atoi(strings.TrimSpace(data))
		nums[i] = n
	}

	ans := fmt.Sprintln(lonelyInteger(nums))
	rw.WriteString(ans)
}

// SOLUTION ==============================================
func lonelyInteger(A []int) (loner int) {
	sort.Ints(A)
	prev, cnt := A[0], 1
	for _, n := range A[1:] {
		if prev != n {
			if cnt > 1 {
				cnt = 0
			} else {
				loner = prev
				return
			}
		}
		prev = n
		cnt++
	}
	loner = A[len(A)-1]
	return
}
